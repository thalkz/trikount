package database

import (
	"github.com/pkg/errors"
	"github.com/thalkz/trikount/models"
)

func AddExpense(projectId string, title string, amount float64, paidBy int, spendBy []int) error {
	row := db.QueryRow(`INSERT INTO expenses (title, amount, project_id, paid_by) 
		values($1, $2, $3, $4) RETURNING id`, title, amount, projectId, paidBy)
	var expenseId int
	err := row.Scan(&expenseId)
	if err != nil {
		return errors.Wrap(err, "failed to scan expense row")
	}

	err = row.Err()
	if err != nil {
		return errors.Wrap(err, "failed to insert expense")
	}

	for _, spenderId := range spendBy {
		_, err = db.Exec(`INSERT INTO spent_by (expense_id, member_id) values($1, $2)`, expenseId, spenderId)
		if err != nil {
			return errors.Wrap(err, "failed to insert spender")
		}
	}

	return nil
}

func GetExpenses(id string) ([]*models.Expense, error) {
	rows, err := db.Query(`SELECT expenses.id, title, amount, members.id, members.name
		FROM expenses 
		JOIN members ON expenses.paid_by = members.id
		WHERE expenses.project_id = $1`, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get expenses")
	}

	var expenses []*models.Expense
	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.PaidBy.Id, &expense.PaidBy.Name)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan expense row")
		}
		expenses = append(expenses, &expense)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan expenses")
	}

	return expenses, nil
}
