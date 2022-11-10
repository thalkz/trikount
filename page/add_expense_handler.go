package page

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddExpense(c *gin.Context) {
	projectId := c.Param("projectId")

	expenseName := c.Query("name")
	expenseAmountStr := c.Query("amount")

	if expenseName == "" || expenseAmountStr == "" {
		c.HTML(http.StatusOK, "add_expense.html", nil)
		return
	}

	expenseAmount, err := strconv.ParseFloat(expenseAmountStr, 32)
	if err != nil {
		c.HTML(http.StatusInternalServerError, err.Error(), nil)
	}

	// TODO Save expense
	_ = expenseAmount
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s", projectId))
}
