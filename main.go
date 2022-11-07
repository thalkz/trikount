package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	page "github.com/thalkz/trikount/pages"
)

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("www/*")

	r.GET("/ping", handlePing)

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

// func main() {
// 	db, err := sql.Open("sqlite3", "trikount.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	sqlStmt := `CREATE TABLE projects (id integer NOT NULL PRIMARY KEY, name text)`
// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		log.Printf("%q: %s\n", err, sqlStmt)
// 	}

// 	tx, err := db.Begin()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	stmt, err := tx.Prepare("INSERT INTO projects(id, name) values(?, ?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()
// 	for i := 0; i < 100; i++ {
// 		_, err = stmt.Exec(i, fmt.Sprintf("project-%03d", i))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	rows, err := db.Query("SELECT id, name FROM projects")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	var id int
// 	var name string
// 	for rows.Next() {
// 		err = rows.Scan(&id, &name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(id, name)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
