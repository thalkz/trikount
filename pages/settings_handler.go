package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Settings(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.html", nil)
}
