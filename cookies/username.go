package cookies

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func GetCurrentUsername(c *gin.Context, projectId string) (string, error) {
	encodedUsername, err := c.Cookie(projectId)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get cookie from key=%v", projectId)
	}

	decodedUsername, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		return "", errors.Wrapf(err, "failed to unescape value %v", encodedUsername)
	}

	return decodedUsername, nil
}
