package database

import (
	"github.com/pkg/errors"
	"github.com/thalkz/trikount/models"
)

func AddMember(projectId string, name string) error {
	_, err := db.Exec(`INSERT INTO members (name, project_id) values($1, $2)`, name, projectId)
	if err != nil {
		return errors.Wrapf(err, "failed to insert member in project %v (name=%v)", projectId, name)
	}
	return nil
}

func GetMembers(projectId string) ([]*models.Member, error) {
	rows, err := db.Query(`SELECT id, name FROM members WHERE project_id = $1`, projectId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get members in project %v", projectId)
	}
	var members []*models.Member
	for rows.Next() {
		var member models.Member
		err := rows.Scan(&member.Id, &member.Name)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to scan member in project %v", projectId)
		}
		members = append(members, &member)
	}
	return members, nil
}

func GetMemberByName(projectId string, name string) (*models.Member, error) {
	row := db.QueryRow(`SELECT id, name FROM members WHERE project_id = $1 AND name = $2`, projectId, name)
	if err := row.Err(); err != nil {
		return nil, errors.Wrapf(err, "failed to get member by name %v", name)
	}
	var member models.Member
	err := row.Scan(&member.Id, &member.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to scan member by name %v", name)
	}
	return &member, nil
}

func GetMemberById(projectId string, id int) (*models.Member, error) {
	row := db.QueryRow(`SELECT id, name FROM members WHERE project_id = $1 AND id = $2`, projectId, id)
	if err := row.Err(); err != nil {
		return nil, errors.Wrapf(err, "failed to get member by id %v", id)
	}
	var member models.Member
	err := row.Scan(&member.Id, &member.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to scan member by id %v", id)
	}
	return &member, nil
}
