package page

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/hash"
)

func CreateProject(c *gin.Context) {
	projectName := c.Query("name")

	if projectName == "" {
		c.HTML(http.StatusOK, "create_project.html", nil)
	} else {
		projectId, err := findAvailableProjectId()
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		err = database.CreateProject(projectId, projectName, time.Now())
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		c.Redirect(http.StatusFound, fmt.Sprintf("/%s", projectId))
	}
}

func findAvailableProjectId() (string, error) {
	maxRetries := 5
	projectIdLength := 6
	for retry := 0; retry < 5; retry++ {
		projectId := hash.NewHash(projectIdLength)
		exists, err := database.CheckExists(projectId)
		if err != nil {
			return "", errors.Wrap(err, "failed to check if project id exists")
		}
		if !exists {
			return projectId, nil
		}
	}
	return "", fmt.Errorf("could not find available project id after %v retries", maxRetries)
}
