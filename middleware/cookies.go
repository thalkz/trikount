package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/cookies"
)

func SetProjectCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := c.Param("projectId")
		if projectId != "" {
			cookies.SaveProjectId(c, projectId)
		}
	}
}

func SetUserCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdStr, exists := c.GetQuery("user_id")
		if exists {
			projectId := c.Param("projectId")
			cookies.SetUserId(c, projectId, userIdStr)
		}
	}
}
