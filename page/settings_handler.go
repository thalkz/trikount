package page

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

func Settings() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := c.Param("projectId")
		deleteProject := c.Query("delete") == "on"
		projectName := c.Query("name")

		if deleteProject {
			handleDeleteProject(c, projectId)
		} else if projectName != "" {
			handleRenameProject(c, projectId, projectName)
		} else {
			renderSettingsPage(c, projectId)
		}
	}
}

func handleDeleteProject(c *gin.Context, projectId string) {
	database.DeleteProject(projectId)
	c.Redirect(http.StatusFound, "/")
}

func handleRenameProject(c *gin.Context, projectId string, projectName string) {
	err := database.RenameProject(projectId, projectName)
	if err != nil {
		error_helper.HTML(http.StatusBadRequest, err, c)
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/t/%s", projectId))
}

func renderSettingsPage(c *gin.Context, projectId string) {
	type page struct {
		Project         *models.Project
		Members         []*models.Member
		CurrentUsername string
	}

	project, err := database.GetProject(projectId)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}

	members, err := database.GetMembers(projectId)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}

	currentUsername, _ := c.Cookie(projectId)

	c.HTML(http.StatusOK, "settings.html", page{
		Project:         project,
		Members:         members,
		CurrentUsername: currentUsername,
	})
}
