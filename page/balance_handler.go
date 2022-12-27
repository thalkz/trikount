package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

func Balance() gin.HandlerFunc {
	type page struct {
		Balance   []*models.MemberBalance
		Transfers []*models.Transfer
	}

	return func(c *gin.Context) {
		projectId := c.Param("projectId")

		balance, err := database.GetBalance(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		transfers := balance.GetTransfers()

		c.HTML(http.StatusOK, "balance.html", page{
			Balance:   balance.Members,
			Transfers: transfers,
		})
	}
}
