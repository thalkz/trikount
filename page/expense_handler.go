package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type expensePage struct {
	Name   string
	Amount float32
	Parts  []expensePagePart
}

type expensePagePart struct {
	MemberName string
	Amount     float32
}

func Expense(c *gin.Context) {
	c.HTML(http.StatusOK, "expense.html", expensePage{
		Name:   "Resto",
		Amount: 11.50,
		Parts: []expensePagePart{
			{
				MemberName: "Roland",
				Amount:     1.0,
			},
			{
				MemberName: "H",
				Amount:     10.5,
			},
		},
	})
}
