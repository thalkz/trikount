package database

import (
	"database/sql"
	"os"

	"github.com/pkg/errors"
)

var db *sql.DB

func Setup() (close func() error, err error) {
	err = os.Mkdir("data", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		errors.Wrap(err, "failed to create folder")
		return
	}

	db, err = sql.Open("sqlite3", "data/data.db")
	if err != nil {
		errors.Wrap(err, "failed to open database")
	}
	close = db.Close

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS projects (
		id text NOT NULL PRIMARY KEY,
		name text NOT NULL,
		created_at text NOT NULL
		)`)
	if err != nil {
		err = errors.Wrap(err, "failed to create projects table")
		return
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS members (
		id integer NOT NULL PRIMARY KEY,
		name text NOT NULL,
		project_id text NOT NULL,
		FOREIGN KEY (project_id) REFERENCES projects(id)
	)`)
	if err != nil {
		err = errors.Wrap(err, "failed to create members table")
		return
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS expenses (
		id integer NOT NULL PRIMARY KEY,
		project_id text,
		title text NOT NULL,
		amount real NOT NULL,
		paid_by integer,
		FOREIGN KEY (project_id) REFERENCES projects(id),
		FOREIGN KEY (paid_by) REFERENCES members(id)
	)`)
	if err != nil {
		err = errors.Wrap(err, "failed to create expenses table")
		return
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS spent_by (
		expense_id integer,
		member_id integer,
		FOREIGN KEY (expense_id) REFERENCES expenses(id),
		FOREIGN KEY (member_id) REFERENCES members(id)
	)`)
	if err != nil {
		err = errors.Wrap(err, "failed to create spent_by table")
		return
	}

	// Create a `parts` view, that is used to query the current balance
	_, err = db.Exec(`CREATE VIEW IF NOT EXISTS v_parts AS
		SELECT expense_id, project_id, title, amount / COUNT(*) amount
			FROM expenses
			JOIN spent_by ON expenses.id = spent_by.expense_id
			GROUP BY expenses.id`)
	if err != nil {
		err = errors.Wrap(err, "failed to create parts view")
		return
	}

	_, err = db.Exec(`CREATE VIEW IF NOT EXISTS v_spent_balance AS
		SELECT spent_by.member_id, SUM(amount) amount
			FROM v_parts
			JOIN spent_by ON spent_by.expense_id = v_parts.expense_id
			GROUP BY spent_by.member_id;`)
	if err != nil {
		err = errors.Wrap(err, "failed to create spent_balance view")
		return
	}

	_, err = db.Exec(`CREATE VIEW IF NOT EXISTS v_paid_balance AS
		SELECT paid_by AS member_id, project_id, SUM(amount) amount 
		FROM expenses 
		GROUP BY paid_by`)
	if err != nil {
		err = errors.Wrap(err, "failed to create paid_balance view")
		return
	}

	return
}
