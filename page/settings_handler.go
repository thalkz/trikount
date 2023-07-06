package page

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thalkz/trikount/cookies"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

func Settings() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := c.Param("projectId")
		deleteProject := c.Query("delete") == "on"
		removeMemberId := c.Query("remove_member")
		projectName := c.Query("name")

		if deleteProject {
			handleDeleteProject(c, projectId)
		} else if removeMemberId != "" {
			handleRemoveMember(c, projectId, removeMemberId)
		} else if projectName != "" {
			handleRenameProject(c, projectId, projectName)
		} else {
			renderSettingsPage(c, projectId)
		}
	}
}

func handleDeleteProject(c *gin.Context, projectId string) {
	err := database.DeleteProject(projectId)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func handleRemoveMember(c *gin.Context, projectId string, removeMemberId string) {
	memberId, err := strconv.Atoi(removeMemberId)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, errors.Wrap(err, "failed to parse member_id to remove"), c)
		return
	}

	err = database.RemoveMember(projectId, memberId)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/t/%s/settings", projectId))
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
		Project *models.Project
		Members []*models.Member
		UserId  int
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

	userId, _ := cookies.GetUserId(c, projectId)

	c.HTML(http.StatusOK, "settings.html", page{
		Project: project,
		Members: members,
		UserId:  userId,
	})
}
