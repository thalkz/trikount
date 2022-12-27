package page

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
)

func AddMember() gin.HandlerFunc {
	return func(c *gin.Context) {

		projectId := c.Param("projectId")
		memberName := c.Query("name")

		if memberName == "" {
			c.HTML(http.StatusOK, "add_member.html", nil)
			return
		}

		err := database.AddMember(projectId, memberName)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		c.Redirect(http.StatusFound, fmt.Sprintf("/t/%s", projectId))
	}
}
