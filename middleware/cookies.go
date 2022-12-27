package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

const cookieExpireSeconds = 60 * 60 * 24 * 365 // 1 Year

func SetProjectCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		current := c.Param("projectId")
		cookie, err := c.Cookie("project_ids")
		if err != nil {
			log.Printf("failed to get cookie: %v", cookie)
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
		log.Printf("DEBUG: setting cookie project_ids=%v", cookie)
		c.SetCookie("project_ids", cookie, cookieExpireSeconds, "/", "", false, true)
	}
}
