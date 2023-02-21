package page

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

func Home() gin.HandlerFunc {
	type page struct {
		Projects []*models.Project
	}

	return func(c *gin.Context) {
		projectIds := []string{}
		projectIdsStr, err := c.Cookie("project_ids")
		if err != nil {
			projectIds = strings.Split(projectIdsStr, ",")
		}

		projects, err := database.GetProjects(projectIds)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		log.Printf("DEBUG: projectIds: %v", projectIds)
		log.Printf("DEBUG: projects: %v", projects)

		c.HTML(http.StatusOK, "home.html", page{
			Projects: projects,
		})
	}
}
