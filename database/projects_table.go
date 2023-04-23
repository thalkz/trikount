package database

import (
	"database/sql"
	"fmt"
	"strings"
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

func DeleteProject(id string) error {
	_, err := db.Exec(`DELETE FROM projects WHERE id = $1`, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete project")
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

func RenameProject(projectId string, name string) error {
	_, err := db.Exec(`UPDATE projects SET name = $1 WHERE id = $2`, name, projectId)
	if err != nil {
		return errors.Wrap(err, "failed to rename project")
	}
	return nil
}

func GetProjects(ids []string) ([]*models.Project, error) {
	arr := fmt.Sprintf("('%v')", strings.Join(ids, "','"))
	stmt := fmt.Sprintf(`SELECT id, name, created_at FROM projects WHERE id IN %v`, arr)
	// TODO This is probably vulnerable to SQL injection. Find a more secure way of encoding the array
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan projects")
	}

	projects := make([]*models.Project, 0)
	for rows.Next() {
		var project models.Project
		var createdAtStr string
		err := rows.Scan(&project.Id, &project.Name, &createdAtStr)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get project")
		}
		project.CreatedAt, err = time.Parse(time.UnixDate, createdAtStr)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse time")
		}
		projects = append(projects, &project)
	}

	return projects, nil
}
