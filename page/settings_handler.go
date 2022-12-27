package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Settings() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "settings.html", nil)
	}
}
