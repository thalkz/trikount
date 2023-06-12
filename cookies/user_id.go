package cookies

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func toCookieKey(projectId string) string {
	return fmt.Sprintf("%v_user_id", projectId)
}

func SetUserId(c *gin.Context, projectId string, userIdStr string) {
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		fmt.Printf("failed to parse userId to integer %v\n", userIdStr)
		return
	}

	cookieExpireSeconds := 60 * 60 * 24 * 365 // 1 Year
	key := toCookieKey(projectId)
	value := fmt.Sprintf("%v", userId)
	c.SetCookie(key, value, cookieExpireSeconds, "/", "", false, true)
}

func GetUserId(c *gin.Context, projectId string) (int, error) {
	key := toCookieKey(projectId)
	userIdStr, err := c.Cookie(key)
	if err != nil {
		return -1, errors.Wrapf(err, "failed to get userId from key=%v", key)
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return -1, errors.Wrapf(err, "failed to parse userId %v", userIdStr)
	}

	return userId, nil
}
