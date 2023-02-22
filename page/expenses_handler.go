package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

func Expenses() gin.HandlerFunc {
	type page struct {
		ProjectId    string
		Name         string
		Expenses     []*models.Expense
		ShowTutorial bool
	}

	return func(c *gin.Context) {
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

		c.HTML(http.StatusOK, "expenses.html", page{
			ProjectId: projectId,
			Name:      project.Name,
			Expenses:  expenses,
		})
	}
}
