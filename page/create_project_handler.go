package page

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	projectName := c.Query("name")

	if projectName == "" {
		c.HTML(http.StatusOK, "create_project.html", nil)
	} else {
		// TODO Register project
		projectId := "PROJECT_ID"
		c.Redirect(http.StatusFound, fmt.Sprintf("/%s", projectId))
	}
}
