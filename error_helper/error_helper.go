package error_helper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorPage struct {
	Code    int
	Title   string
	Message string
}

func HTML(code int, err error, c *gin.Context) {
	fmt.Printf("ERROR: [%v] %v (at %v) %v\n", code, http.StatusText(code), c.Request.URL.Path, err)

	e := errorPage{
		Code:    code,
		Title:   http.StatusText(code),
		Message: err.Error(),
	}

	c.HTML(code, "error.html", e)
}
