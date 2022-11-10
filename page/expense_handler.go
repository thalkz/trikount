package page

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

type expensePage struct {
	Expense *models.Expense
}

func Expense(c *gin.Context) {
	expenseIdStr := c.Param("expenseId")

	expenseId, err := strconv.Atoi(expenseIdStr)
	if err != nil {
		error_helper.HTML(http.StatusBadRequest, err, c)
	}

	expense, err := database.GetExpense(expenseId)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
	}

	c.HTML(http.StatusOK, "expense.html", expensePage{
		Expense: expense,
	})
}
