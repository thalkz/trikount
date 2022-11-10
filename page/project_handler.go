package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

type projectPage struct {
	ProjectId string
	Name      string
	Expenses  []*models.Expense
}

func Project(c *gin.Context) {
	projectId := c.Param("projectId")

	project, err := database.GetProject(projectId)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}

	expenses, err := database.GetExpenses(projectId)
	if err != nil {
		error_helper.HTML(http.StatusInternalServerError, err, c)
		return
	}

	data := projectPage{
		ProjectId: projectId,
		Name:      project.Name,
		Expenses:  expenses,
	}

	c.HTML(http.StatusOK, "project.html", data)
}
