package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/cookies"
)

const cookieExpireSeconds = 60 * 60 * 24 * 365 // 1 Year

func SetProjectCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		current := c.Param("projectId")
		cookie, err := c.Cookie("project_ids")
		if err != nil {
			log.Printf("ERROR: Failed to get cookie: %v", cookie)
		}
		set := make(map[string]bool)
		set[current] = true
		for _, id := range strings.Split(cookie, ",") {
			if id != "" {
				set[id] = true
			}
		}
		cleaned := make([]string, 0)
		for id := range set {
			cleaned = append(cleaned, id)
		}
		cookie = strings.Join(cleaned, ",")
		log.Printf("DEBUG: Setting cookie project_ids=%v\n", cookie)
		c.SetCookie("project_ids", cookie, cookieExpireSeconds, "/", "", false, true)
	}
}

func SetUserIdCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdStr, exists := c.GetQuery("user_id")
		if exists {
			projectId := c.Param("projectId")
			cookies.SetUserId(c, projectId, userIdStr)
		}
	}
}
