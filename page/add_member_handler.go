package page

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
)

func AddMembers() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := c.Param("projectId")
		memberNames := c.QueryArray("name")

		if len(memberNames) > 0 {
			handleAddMembers(c, projectId, memberNames)
		} else {
			renderAddMembersPage(c)
		}
	}
}

func handleAddMembers(c *gin.Context, projectId string, memberNames []string) {
	for _, name := range memberNames {
		err := database.AddMember(projectId, name)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/t/%s", projectId))
}

func renderAddMembersPage(c *gin.Context) {
	c.HTML(http.StatusOK, "add_members.html", nil)
}
