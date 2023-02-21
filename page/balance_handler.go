package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/format"
	"github.com/thalkz/trikount/models"
)

func Balance() gin.HandlerFunc {
	type page struct {
		ProjectId  string
		Balance    []*models.MemberBalance
		TotalSpent string
		Transfers  []*models.Transfer
	}

	return func(c *gin.Context) {
		projectId := c.Param("projectId")

		balance, err := database.GetBalance(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		totalSpent, err := database.GetTotalSpent(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		c.HTML(http.StatusOK, "balance.html", page{
			ProjectId:  projectId,
			Transfers:  balance.GetTransfers(),
			TotalSpent: format.ToEuro(totalSpent),
			Balance:    balance.Members,
		})
	}
}
