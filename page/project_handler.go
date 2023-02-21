package page

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/models"
)

func Project() gin.HandlerFunc {
	type page struct {
		ProjectId    string
		Name         string
		Expenses     []*models.Expense
		ShowTutorial bool
	}

	return func(c *gin.Context) {
		projectId := c.Param("projectId")
		showTutorial := shouldShowTutorial(c)

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

		c.HTML(http.StatusOK, "project.html", page{
			ProjectId:    projectId,
			Name:         project.Name,
			Expenses:     expenses,
			ShowTutorial: showTutorial,
		})
	}
}

func shouldShowTutorial(c *gin.Context) bool {
	if c.Query("show_tutorial") == "false" {
		log.Printf("DEBUG: Set show_tutorial=false")
		c.SetCookie("show_tutorial", "false", 0, "/", "", false, true)
		return false
	} else {
		showTutorial, _ := c.Cookie("show_tutorial")
		return showTutorial == "true"
	}
}
