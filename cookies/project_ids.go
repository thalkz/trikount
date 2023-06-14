package cookies

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const cookieExpireSeconds = 60 * 60 * 24 * 365 // 1 Year

func GetSavedProjectIds(c *gin.Context) []string {
	projectIds := []string{}
	projectIdsStr, err := c.Cookie("project_ids")
	if err == nil {
		projectIds = strings.Split(projectIdsStr, ",")
	}
	return projectIds
}

func SaveProjectId(c *gin.Context, current string) {
	array := GetSavedProjectIds(c)
	ids := make(map[string]bool)
	ids[current] = true
	for _, id := range array {
		if id != "" {
			ids[id] = true
		}
	}
	cleaned := make([]string, 0)
	for id := range ids {
		cleaned = append(cleaned, id)
	}
	result := strings.Join(cleaned, ",")
	c.SetCookie("project_ids", result, cookieExpireSeconds, "/", "", false, true)
}
