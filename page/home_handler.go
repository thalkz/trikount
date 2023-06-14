package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/cookies"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

func Home() gin.HandlerFunc {
	type page struct {
		Projects []*models.Project
	}

	return func(c *gin.Context) {
		projectIds := cookies.GetSavedProjectIds(c)

		projects, err := database.GetProjects(projectIds)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		c.HTML(http.StatusOK, "home.html", page{
			Projects: projects,
		})
	}
}
