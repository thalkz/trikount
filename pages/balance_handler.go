package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type balancePage struct {
	ProjectName string
	Transfers   []balancePageTransfer
}

type balancePageTransfer struct {
	Amount float32
	From   string
	To     string
}

func Balance(c *gin.Context) {
	c.HTML(http.StatusOK, "balance.html", balancePage{
		ProjectName: "Project Name",
		Transfers: []balancePageTransfer{
			{
				Amount: 100.0,
				From:   "Roland",
				To:     "H",
			},
		},
	})
}
