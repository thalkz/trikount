package page

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddMember(c *gin.Context) {
	projectId := c.Param("projectId")
	memberName := c.Query("name")

	if memberName == "" {
		c.HTML(http.StatusOK, "add_member.html", nil)
		return
	}

	// TODO Add Member
	c.Redirect(http.StatusFound, fmt.Sprintf("/%s", projectId))
}
