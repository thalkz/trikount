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

func UnsetUserId(c *gin.Context, projectId string) {
	key := toCookieKey(projectId)
	c.SetCookie(key, "", 0, "/", "", false, true)
}

func GetUserId(c *gin.Context, projectId string) (int, error) {
	userIdStr, exists := c.GetQuery("user_id")
	if exists && userIdStr == "" {
		return -1, errors.Errorf("user_id is being unset")
	} else if exists {
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			return -1, errors.Wrapf(err, "failed to parse user_id %v", userIdStr)
		}
		return userId, nil
	}

	key := toCookieKey(projectId)
	cookieUserIdStr, err := c.Cookie(key)
	if err != nil {
		return -1, errors.Wrapf(err, "failed to get userId from key=%v", key)
	}

	userId, err := strconv.Atoi(cookieUserIdStr)
	if err != nil {
		return -1, errors.Wrapf(err, "failed to parse userId %v", cookieUserIdStr)
	}

	return userId, nil
}
