package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

type balancePage struct {
	Balance   []*models.MemberBalance
	Transfers []*models.Transfer
}

func Balance(c *gin.Context) {
	projectId := c.Param("projectId")

	balance, err := database.GetBalance(projectId)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}

	transfers := balance.GetTransfers()

	c.HTML(http.StatusOK, "balance.html", balancePage{
		Balance:   balance.Members,
		Transfers: transfers,
	})
}
