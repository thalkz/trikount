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
			handleDeleteExpense(c, projectId, expenseId)
		} else {
			handleRenderExpensePage(c, projectId, expenseId)
		}
	}
}

func handleDeleteExpense(c *gin.Context, projectId string, expenseId int) {
	err := database.DeleteExpense(projectId, expenseId)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/t/%s/expenses", projectId))
}

func handleRenderExpensePage(c *gin.Context, projectId string, expenseId int) {
	type page struct {
		ProjectId string
		Expense   *models.Expense
	}

	expense, err := database.GetExpense(projectId, expenseId)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}

	c.HTML(http.StatusOK, "expense.html", page{
		Expense:   expense,
		ProjectId: projectId,
	})
}
