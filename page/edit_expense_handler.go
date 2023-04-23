package page

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

func EditExpense() gin.HandlerFunc {
	return func(c *gin.Context) {
		expenseIdStr := c.Param("expenseId")
		projectId := c.Param("projectId")
		title := c.Query("name")

		expenseId, err := strconv.Atoi(expenseIdStr)
		if err != nil {
			error_helper.HTML(http.StatusBadRequest, err, c)
			return
		}

		members, err := database.GetMembers(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		if title == "" {
			handleRenderEditExpensePage(c, expenseId, members)
		} else {
			handleEditExpense(c, projectId, members, expenseId, title)
		}
	}
}

func handleEditExpense(c *gin.Context, projectId string, members []*models.Member, expenseId int, title string) {
	amountStr := c.Query("amount")
	paidByStr := c.Query("paid_by")
	isTransfer := c.Query("isTransfer") == "on"

	amount, err := strconv.ParseFloat(amountStr, 32)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}
	paidBy, err := strconv.Atoi(paidByStr)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}

	var spendBy []int
	for _, member := range members {
		if c.Query(fmt.Sprintf("%v", member.Id)) == "on" {
			spendBy = append(spendBy, member.Id)
		}
	}

	now := time.Now()
	err = database.EditExpense(projectId, expenseId, title, amount, paidBy, spendBy, isTransfer, now)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/t/%s", projectId))
}

func handleRenderEditExpensePage(c *gin.Context, expenseId int, members []*models.Member) {
	type page struct {
		IsEdit  bool
		Members []*models.Member
		Expense *models.Expense
	}

	expense, err := database.GetExpense(expenseId)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}

	// Order members so that the paidBy member is first in the slice
	for i, member := range members {
		if member.Id == expense.PaidBy.Id {
			tmp := members[0]
			members[0] = members[i]
			members[i] = tmp
		}
	}

	c.HTML(http.StatusOK, "edit_expense.html", page{
		IsEdit:  true,
		Expense: expense,
		Members: members,
	})
}
