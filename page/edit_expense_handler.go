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

func EditExpense() gin.HandlerFunc {
	type page struct {
		IsEdit  bool
		Members []*models.Member
		Expense *models.Expense
	}

	return func(c *gin.Context) {
		expenseIdStr := c.Param("expenseId")
		projectId := c.Param("projectId")
		title := c.Query("name")
		isTransferStr := c.Query("isTransfer")

		expenseId, err := strconv.Atoi(expenseIdStr)
		if err != nil {
			error_helper.HTML(http.StatusBadRequest, err, c)
		}

		members, err := database.GetMembers(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
		}

		if title == "" {
			expense, err := database.GetExpense(expenseId)
			if err != nil {
				error_helper.HTML(http.StatusInternalServerError, err, c)
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
			return
		}

		// Edit expense
		amountStr := c.Query("amount")
		paidByStr := c.Query("paid_by")

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

		isTransfer := isTransferStr == "on"

		err = database.EditExpense(projectId, expenseId, title, amount, paidBy, spendBy, isTransfer)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		c.Redirect(http.StatusFound, fmt.Sprintf("/t/%s", projectId))
	}
}
