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
	type page struct {
		IsEdit  bool
		Members []*models.Member
		Expense *models.Expense
	}

	return func(c *gin.Context) {
		projectId := c.Param("projectId")

		save := c.Query("save")
		title := c.Query("name")
		amountStr := c.Query("amount")
		paidByStr := c.Query("paid_by")
		isTransferStr := c.Query("is_transfer")

		amount, _ := strconv.ParseFloat(amountStr, 32)
		paidBy, _ := strconv.Atoi(paidByStr)
		members, err := database.GetMembers(projectId)
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

		isTransfer := isTransferStr == "on"

		if save != "on" {
			var paidByMember models.Member
			var spendByMembers []models.Member
			for _, member := range members {
				if member.Id == paidBy {
					paidByMember = *member
				}
				for _, id := range spendBy {
					if member.Id == id {
						spendByMembers = append(spendByMembers, *member)
					}
				}
			}

			c.HTML(http.StatusOK, "edit_expense.html", page{
				IsEdit: false,
				Expense: &models.Expense{
					Title:      title,
					Amount:     amount,
					PaidBy:     paidByMember,
					SpentBy:    spendByMembers,
					IsTransfer: isTransfer,
				},
				Members: members,
			})
			return
		} else {
			now := time.Now()
			err = database.AddExpense(projectId, title, amount, paidBy, spendBy, isTransfer, now)
			if err != nil {
				error_helper.HTML(http.StatusInternalServerError, err, c)
				return
			}
			c.Redirect(http.StatusFound, fmt.Sprintf("/t/%s", projectId))
		}
	}
}
