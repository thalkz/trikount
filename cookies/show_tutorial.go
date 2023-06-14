package cookies

import (
	"github.com/gin-gonic/gin"
)

func SetShowTutorial(c *gin.Context) {
	const cookieExpireSeconds = 24 * 60 * 60 // 1 day
	c.SetCookie("show_tutorial", "true", cookieExpireSeconds, "/", "", false, true)
}

func UnsetShowTutorial(c *gin.Context) {
	c.SetCookie("show_tutorial", "false", 0, "/", "", false, true)
}

func ShouldShowTutorial(c *gin.Context) bool {
	showTutorial, _ := c.Cookie("show_tutorial")
	return showTutorial == "true"
}
