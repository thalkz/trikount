package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

func Settings() gin.HandlerFunc {
	type page struct {
		Project *models.Project
	}

	return func(c *gin.Context) {
		projectId := c.Param("projectId")
		deleteProject := c.Query("delete") == "true"

		if deleteProject {
			database.DeleteProject(projectId)
			c.Redirect(http.StatusFound, "/")
		}

		project, err := database.GetProject(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		c.HTML(http.StatusOK, "settings.html", page{
			Project: project,
		})
	}
}
