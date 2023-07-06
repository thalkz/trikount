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

func RemoveMember(projectId string, id int) error {
	isActive := isMemberActive(id)
	if isActive {
		return errors.Errorf("failed to remove member with id = %v: member is part of one or more expenses", id)
	}

	_, err := db.Exec(`DELETE FROM members WHERE project_id = $1 AND id = $2`, projectId, id)
	if err != nil {
		return errors.Wrapf(err, "failed to remove member in project %v (id=%v)", projectId, id)
	}

	return nil
}

func isMemberActive(userId int) bool {
	var isActive bool
	rows, err := db.Query(`SELECT * FROM spent_by WHERE member_id = $1`, userId)
	if err != nil {
		isActive = true
	}

	for rows.Next() {
		isActive = true
	}

	rows, err = db.Query(`SELECT * FROM expenses WHERE paid_by = $1`, userId)
	if err != nil {
		isActive = true
	}
	for rows.Next() {
		isActive = true
	}

	return isActive
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
