package page

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/error_helper"
)

func Home(c *gin.Context) {
	projectIds, err := c.Cookie("project_ids")
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}

	log.Printf("projectIds %v", projectIds)

	c.HTML(http.StatusOK, "home.html", nil)
}
