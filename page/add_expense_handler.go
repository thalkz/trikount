package page

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thalkz/trikount/cookies"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

func AddExpense() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := c.Param("projectId")
		members, err := database.GetMembers(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		save := c.Query("save")
		title := c.Query("name")
		isTransfer := c.Query("is_transfer") == "on"

		var spentBy []int
		for _, member := range members {
			if c.Query(fmt.Sprintf("%v", member.Id)) == "on" {
				spentBy = append(spentBy, member.Id)
			}
		}

		if save == "on" {
			handleAddExpense(c, projectId, title, members, spentBy, isTransfer)
		} else {
			handleRenderAddExpensePage(c, projectId, title, members, spentBy, isTransfer)
		}
	}
}

func handleAddExpense(c *gin.Context, projectId string, title string, members []*models.Member, spentBy []int, isTransfer bool) {
	amountStr := c.Query("amount")
	amount, err := strconv.ParseFloat(amountStr, 32)
	if err != nil {
		err = errors.Wrap(err, "failed to parse amount")
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	} else if amount <= 0 {
		err = fmt.Errorf("cannot add an expense with amount <= 0")
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}

	paidByStr := c.Query("paid_by")
	paidBy, err := strconv.Atoi(paidByStr)
	if err != nil {
		err = errors.Wrap(err, "failed to parse paid_by")
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}
	isPaidByMember := false
	for _, member := range members {
		if member.Id == paidBy {
			isPaidByMember = true
			break
		}
	}
	if !isPaidByMember {
		err = errors.Wrap(err, "paid_by is not a member of the project")
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}

	now := time.Now()
	err = database.AddExpense(projectId, title, amount, paidBy, spentBy, isTransfer, now)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/t/%s", projectId))
}

func handleRenderAddExpensePage(c *gin.Context, projectId string, title string, members []*models.Member, spentBy []int, isTransfer bool) {
	type page struct {
		IsEdit  bool
		Members []*models.Member
		Expense *models.Expense
	}

	amountStr := c.Query("amount")
	amount, _ := strconv.ParseFloat(amountStr, 32)

	paidByStr := c.Query("paid_by")
	paidById, _ := strconv.Atoi(paidByStr)
	userId, _ := cookies.GetUserId(c, projectId)
	paidByMember := &models.Member{}
	if paidById > 0 {
		var err error
		paidByMember, err = database.GetMemberById(projectId, paidById)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}
	} else if userId != -1 {
		var err error
		paidByMember, err = database.GetMemberById(projectId, userId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}
	}

	var spendByMembers []*models.Member
	if len(spentBy) == 0 {
		spendByMembers = members
	} else {
		for _, member := range members {
			for _, id := range spentBy {
				if id == member.Id {
					spendByMembers = append(spendByMembers, member)
					break
				}
			}
		}
	}

	c.HTML(http.StatusOK, "edit_expense.html", page{
		IsEdit:  false,
		Members: members,
		Expense: &models.Expense{
			Title:      title,
			Amount:     amount,
			PaidBy:     *paidByMember,
			SpentBy:    spendByMembers,
			IsTransfer: isTransfer,
		},
	})
}
