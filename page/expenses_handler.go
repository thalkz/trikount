package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/cookies"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

func Expenses() gin.HandlerFunc {
	type page struct {
		ProjectId    string
		Name         string
		ExpenseParts []*models.ExpenseWithPart
		UserId       int
	}

	return func(c *gin.Context) {
		projectId := c.Param("projectId")
		userId, _ := cookies.GetUserId(c, projectId)

		project, err := database.GetProject(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		expenseParts, err := database.GetExpenseWithParts(projectId, userId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		c.HTML(http.StatusOK, "expenses.html", page{
			ProjectId:    projectId,
			Name:         project.Name,
			ExpenseParts: expenseParts,
			UserId:       userId,
		})
	}
}
