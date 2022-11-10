package database

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

func CreateProject(id, name string, createdAt time.Time) error {
	_, err := db.Exec(`INSERT INTO projects (id, name, created_at) values($1, $2, $3)`, id, name, createdAt)
	if err != nil {
		return errors.Wrap(err, "failed to insert project")
	}
	return nil
}

func CheckExists(id string) (bool, error) {
	err := db.QueryRow(`SELECT id FROM projects WHERE id = $1`, id).Scan()
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, errors.Wrap(err, "failed to check project exists")
	}
	return true, nil
}
