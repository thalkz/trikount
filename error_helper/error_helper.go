package error_helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorPage struct {
	Code    int
	Title   string
	Message string
}

func HTML(code int, err error, c *gin.Context) {
	e := errorPage{
		Code:    code,
		Title:   http.StatusText(code),
		Message: err.Error(),
	}

	c.HTML(code, "error.html", e)
}
