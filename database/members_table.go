package database

import (
	"github.com/pkg/errors"
	"github.com/thalkz/trikount/models"
)

func AddMember(projectId string, name string) error {
	_, err := db.Exec(`INSERT INTO members (name, project_id) values($1, $2)`, name, projectId)
	if err != nil {
		return errors.Wrap(err, "failed to insert member")
	}
	return nil
}

func GetMembers(projectId string) ([]*models.Member, error) {
	rows, err := db.Query(`SELECT id, name FROM members WHERE project_id = $1`, projectId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get members")
	}
	var members []*models.Member
	for rows.Next() {
		var member models.Member
		err := rows.Scan(&member.Id, &member.Name)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan member")
		}
		members = append(members, &member)
	}
	return members, nil
}
