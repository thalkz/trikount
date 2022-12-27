package page

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/error_helper"
)

func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectIdsStr, err := c.Cookie("project_ids")
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		projectIds := strings.Split(projectIdsStr, ",")
		log.Printf("current projects %v", projectIds)

		c.HTML(http.StatusOK, "home.html", nil)
	}
}
