package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

var db *sql.DB

func Setup() (close func() error, err error) {
	db, err = sql.Open("sqlite3", "data.db")
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
		title text NOT NULL,
		project_id text,
		paid_by integer,
		FOREIGN KEY (project_id) REFERENCES projects(id),
		FOREIGN KEY (paid_by) REFERENCES members(id)
	)`)
	if err != nil {
		err = errors.Wrap(err, "failed to create expenses table")
		return
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS expenses_members (
		expense_id integer,
		member_id integer,
		FOREIGN KEY (expense_id) REFERENCES expenses(id),
		FOREIGN KEY (member_id) REFERENCES members(id)
	)`)
	if err != nil {
		err = errors.Wrap(err, "failed to create expenses_members table")
		return
	}

	return
}
