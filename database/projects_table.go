package database

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/thalkz/trikount/models"
)

func CreateProject(id, name string, createdAt time.Time) error {
	_, err := db.Exec(`INSERT INTO projects (id, name, created_at) values($1, $2, $3)`, id, name, createdAt.Format(time.UnixDate))
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

func GetProject(id string) (*models.Project, error) {
	row := db.QueryRow(`SELECT id, name, created_at FROM projects WHERE id = $1`, id)
	var project models.Project
	var createdAtStr string
	err := row.Scan(&project.Id, &project.Name, &createdAtStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get project")
	}
	project.CreatedAt, err = time.Parse(time.UnixDate, createdAtStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse time")
	}

	return &project, nil
}
