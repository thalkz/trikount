package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/middleware"
	"github.com/thalkz/trikount/page"
)

func main() {
	close, err := database.Setup()
	defer close()
	if err != nil {
		log.Fatalf("failed to setup database: %v", err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("www/*")
	r.Static("/assets", "./assets")

	r.GET("/", page.Home)
	r.GET("/create", page.CreateProject)

	project := r.Group("/t")
	project.Use(middleware.SetProjectCookie())
	{
		project.GET("/:projectId/", page.Project)
		project.GET("/:projectId/expenses/add", page.AddExpense)
		project.GET("/:projectId/expenses/:expenseId", page.Expense)
		project.GET("/:projectId/expenses/:expenseId/edit", page.EditExpense)
		project.GET("/:projectId/balance", page.Balance)
		project.GET("/:projectId/members/add", page.AddMember)
		project.GET("/:projectId/settings", page.Settings)
	}

	r.Run()
}
