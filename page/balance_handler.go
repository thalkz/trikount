package page

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/error_helper"
	"github.com/thalkz/trikount/format"
	"github.com/thalkz/trikount/models"
)

func Balance() gin.HandlerFunc {
	type page struct {
		ProjectId      string
		Project        *models.Project
		Balance        []*models.MemberBalance
		TotalSpent     string
		Transfers      []*models.Transfer
		ShowTutorial   bool
		ChooseUsername bool
	}

	return func(c *gin.Context) {
		projectId := c.Param("projectId")
		showTutorial := shouldShowTutorial(c)

		project, err := database.GetProject(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		balance, err := database.GetBalance(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		chooseUsername := shouldChooseUsername(c, projectId, balance.Members)

		totalSpent, err := database.GetTotalSpent(projectId)
		if err != nil {
			error_helper.HTML(http.StatusInternalServerError, err, c)
			return
		}

		c.HTML(http.StatusOK, "balance.html", page{
			ProjectId:      projectId,
			Project:        project,
			Transfers:      balance.GetTransfers(),
			TotalSpent:     format.ToEuro(totalSpent),
			Balance:        balance.Members,
			ShowTutorial:   showTutorial,
			ChooseUsername: chooseUsername,
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

func shouldChooseUsername(c *gin.Context, projectId string, members []*models.MemberBalance) bool {
	username, exists := c.GetQuery("current_username")
	if exists && username == "" {
		return true
	}

	for _, member := range members {
		if member.Name == username {
			return false
		}
	}

	name, err := c.Cookie(projectId)
	if err != nil {
		log.Printf("DEBUG: failed to get cookie %s: %s", projectId, err)
		return true
	}
	for _, member := range members {
		if member.Name == name {
			return false
		}
	}
	return true
}
