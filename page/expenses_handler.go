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
	type item struct {
		Date    string
		Expense *models.ExpenseWithPart
	}

	type page struct {
		ProjectId string
		Name      string
		Content   []item
		UserId    int
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

		content := []item{}
		last := ""
		for _, expense := range expenseParts {
			date := expense.FormattedTimeAgo()
			if last != date {
				last = date
				dateItem := item{
					Date: date,
				}
				content = append(content, dateItem)
			}
			expenseItem := item{
				Expense: expense,
			}
			content = append(content, expenseItem)
		}

		c.HTML(http.StatusOK, "expenses.html", page{
			ProjectId: projectId,
			Name:      project.Name,
			Content:   content,
			UserId:    userId,
		})
	}
}
