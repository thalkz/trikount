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

		members, err := database.GetMembers(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		if save == "on" {
			handleAddExpense(c, projectId, members)
		} else {
			handleRenderAddExpensePage(c, projectId, members)
		}
	}
}

func handleAddExpense(c *gin.Context, projectId string, members []*models.Member) {
	title := c.Query("name")
	amountStr := c.Query("amount")
	amount, _ := strconv.ParseFloat(amountStr, 32)
	paidByStr := c.Query("paid_by")
	paidBy, _ := strconv.Atoi(paidByStr)
	isTransfer := c.Query("is_transfer") == "on"

	var spentBy []int
	for _, member := range members {
		if c.Query(fmt.Sprintf("%v", member.Id)) == "on" {
			spentBy = append(spentBy, member.Id)
		}
	}

	now := time.Now()
	err := database.AddExpense(projectId, title, amount, paidBy, spentBy, isTransfer, now)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/t/%s", projectId))
}

func handleRenderAddExpensePage(c *gin.Context, projectId string, members []*models.Member) {
	type page struct {
		IsEdit  bool
		Members []*models.Member
		Expense *models.Expense
	}

	c.HTML(http.StatusOK, "edit_expense.html", page{
		IsEdit: false,
		Expense: &models.Expense{
			Title:      "",
			Amount:     0,
			PaidBy:     models.Member{},
			SpentBy:    members,
			IsTransfer: false,
		},
		Members: members,
	})
}
