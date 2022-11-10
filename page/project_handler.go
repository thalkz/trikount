package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type projectPage struct {
	ProjectId string
	Name      string
	Expenses  []expense
}

type expense struct {
	Id     int
	Name   string
	Amount int
}

func Project(c *gin.Context) {
	projectId := c.Param("projectId")

	c.HTML(http.StatusOK, "project.html", projectPage{
		ProjectId: projectId,
		Name:      "Project Name",
		Expenses: []expense{
			{
				Id:     1,
				Name:   "Resto",
				Amount: 10,
			},
		},
	})
}
