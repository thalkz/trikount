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

func Expense() gin.HandlerFunc {
	type page struct {
		ProjectId string
		Expense   *models.Expense
	}

	return func(c *gin.Context) {
		expenseIdStr := c.Param("expenseId")
		projectId := c.Param("projectId")
		delete := c.Query("delete")

		expenseId, err := strconv.Atoi(expenseIdStr)
		if err != nil {
			error_helper.HTML(http.StatusBadRequest, err, c)
			return
		}

		if delete == "on" {
			err = database.DeleteExpense(expenseId)
			if err != nil {
				error_helper.HTML(http.StatusInternalServerError, err, c)
				return
			}

			c.Redirect(http.StatusFound, fmt.Sprintf("/t/%s/expenses", projectId))
			return
		}

		expense, err := database.GetExpense(expenseId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		c.HTML(http.StatusOK, "expense.html", page{
			Expense:   expense,
			ProjectId: projectId,
		})
	}
}
