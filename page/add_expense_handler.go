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

func AddExpense() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := c.Param("projectId")
		save := c.Query("save")
		title := c.Query("name")
		amountStr := c.Query("amount")
		paidByStr := c.Query("paid_by")
		isTransfer := c.Query("is_transfer") == "on"
		amount, _ := strconv.ParseFloat(amountStr, 32)
		paidBy, _ := strconv.Atoi(paidByStr)

		members, err := database.GetMembers(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		var spentBy []int
		for _, member := range members {
			if c.Query(fmt.Sprintf("%v", member.Id)) == "on" {
				spentBy = append(spentBy, member.Id)
			}
		}

		if save == "on" {
			handleAddExpense(c, projectId, title, amount, paidBy, spentBy, isTransfer)
		} else {
			handleRenderAddExpensePage(c, projectId, title, amount, paidBy, spentBy, isTransfer, members)
		}
	}
}

func handleAddExpense(c *gin.Context, projectId string, title string, amount float64, paidBy int, spentBy []int, isTransfer bool) {
	now := time.Now()
	err := database.AddExpense(projectId, title, amount, paidBy, spentBy, isTransfer, now)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/t/%s", projectId))
}

func handleRenderAddExpensePage(c *gin.Context, projectId string, title string, amount float64, paidBy int, spentBy []int, isTransfer bool, members []*models.Member) {
	type page struct {
		IsEdit  bool
		Members []*models.Member
		Expense *models.Expense
	}

	var paidByMember models.Member
	var spentByMembers []models.Member
	for _, member := range members {
		if member.Id == paidBy {
			paidByMember = *member
		}
		for _, id := range spentBy {
			if member.Id == id {
				spentByMembers = append(spentByMembers, *member)
			}
		}
	}

	c.HTML(http.StatusOK, "edit_expense.html", page{
		IsEdit: false,
		Expense: &models.Expense{
			Title:      title,
			Amount:     amount,
			PaidBy:     paidByMember,
			SpentBy:    spentByMembers,
			IsTransfer: isTransfer,
		},
		Members: members,
	})
}
