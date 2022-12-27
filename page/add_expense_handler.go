package page

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

func AddExpense() gin.HandlerFunc {
	type page struct {
		IsEdit  bool
		Members []*models.Member
		Expense *models.Expense
	}

	return func(c *gin.Context) {
		projectId := c.Param("projectId")
		title := c.Query("name")
		amountStr := c.Query("amount")
		paidByStr := c.Query("paid_by")

		members, err := database.GetMembers(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
		}

		if title == "" || amountStr == "" {
			c.HTML(http.StatusOK, "edit_expense.html", page{
				IsEdit:  false,
				Expense: &models.Expense{},
				Members: members,
			})
			return
		}

		amount, err := strconv.ParseFloat(amountStr, 32)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
		}
		paidBy, err := strconv.Atoi(paidByStr)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
		}

		var spendBy []int
		for _, member := range members {
			if c.Query(fmt.Sprintf("%v", member.Id)) == "on" {
				spendBy = append(spendBy, member.Id)
			}
		}

		err = database.AddExpense(projectId, title, amount, paidBy, spendBy)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		c.Redirect(http.StatusFound, fmt.Sprintf("/t/%s", projectId))
	}
}
