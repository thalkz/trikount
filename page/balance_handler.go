package page

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thalkz/trikount/cookies"
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
		UserId         int
	}

	return func(c *gin.Context) {
		projectId := c.Param("projectId")
		hideTutorial := c.Query("show_tutorial") == "false"

		if hideTutorial {
			cookies.UnsetShowTutorial(c)
		}

		shouldShowTutorial := !hideTutorial && cookies.ShouldShowTutorial(c)

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

		userId, err := cookies.GetUserId(c, projectId)
		fmt.Printf("userId %v, err=%v", userId, err)
		chooseUsername := err != nil

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
			ShowTutorial:   shouldShowTutorial,
			ChooseUsername: chooseUsername,
			UserId:         userId,
		})
	}
}
