package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"github.com/thalkz/trikount/database"
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
	r.GET("/:projectId", page.Project)
	r.GET("/:projectId/expenses/add", page.AddExpense)
	r.GET("/:projectId/expenses/:expenseId", page.Expense)
	r.GET("/:projectId/balance", page.Balance)
	r.GET("/:projectId/members/add", page.AddMember)
	r.GET("/:projectId/settings", page.Settings)

	r.Run()
}
