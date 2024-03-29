package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"

	"github.com/thalkz/trikount/database"
	"github.com/thalkz/trikount/middleware"
	"github.com/thalkz/trikount/page"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("WARN: No .env file to load, using default envs")
	}

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	close, err := database.Setup()
	defer close()
	if err != nil {
		log.Fatalf("FATAL: failed to setup database: %v", err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("www/*")
	r.Static("/assets", "./assets")

	r.GET("/", page.Home())
	r.GET("/create", page.CreateProject())

	project := r.Group("/t")
	project.Use(middleware.SetProjectCookie())
	project.Use(middleware.SetUserCookie())
	{
		project.GET("/:projectId/", page.Balance())
		project.GET("/:projectId/expenses", page.Expenses())
		project.GET("/:projectId/expenses/add", page.AddExpense())
		project.GET("/:projectId/expenses/:expenseId", page.Expense())
		project.GET("/:projectId/expenses/:expenseId/edit", page.EditExpense())
		project.GET("/:projectId/members/add", page.AddMembers())
		project.GET("/:projectId/settings", page.Settings())
	}

	port := os.Getenv("PORT")
	log.Printf("INFO: server listening on port %v", port)
	r.Run(":" + port)
}
