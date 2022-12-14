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

func EditExpense(projectId string, expenseId int, title string, amount float64, paidBy int, spentBy []int) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin tx")
	}
	defer tx.Rollback()

	_, err = tx.Exec(`UPDATE expenses SET title = ?, amount = ?, paid_by = ? WHERE id = ?`, title, amount, paidBy, expenseId)
	if err != nil {
		return errors.Wrap(err, "failed to update expense")
	}

	_, err = tx.Exec(`DELETE FROM spent_by WHERE expense_id = $1`, expenseId)
	if err != nil {
		return errors.Wrap(err, "failed to delete old spent_by")
	}

	for _, spenderId := range spentBy {
		_, err = tx.Exec(`INSERT INTO spent_by (expense_id, member_id) values($1, $2)`, expenseId, spenderId)
		if err != nil {
			return errors.Wrap(err, "failed to insert spender")
		}
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit tx")
	}

	return nil
}

func GetExpenses(projectId string) ([]*models.Expense, error) {
	rows, err := db.Query(`SELECT expenses.id, title, amount, members.id, members.name
		FROM expenses 
		JOIN members ON expenses.paid_by = members.id
		WHERE expenses.project_id = $1`, projectId)
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

func GetExpense(id int) (*models.Expense, error) {
	row := db.QueryRow(`SELECT expenses.id, expenses.title, expenses.amount, members.id, members.name
		FROM expenses 
		JOIN members ON expenses.paid_by = members.id
		WHERE expenses.id = $1`, id)

	var expense models.Expense
	err := row.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.PaidBy.Id, &expense.PaidBy.Name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan expense")
	}

	err = row.Err()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get expense")
	}

	rows, err := db.Query(`SELECT members.id, members.name 
		FROM spent_by
		JOIN members ON spent_by.member_id = members.id
		WHERE spent_by.expense_id = $1`, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get spent_by members")
	}

	for rows.Next() {
		var member models.Member
		err = rows.Scan(&member.Id, &member.Name)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan spent_by member")
		}
		expense.SpentBy = append(expense.SpentBy, member)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get spent_by members")
	}

	return &expense, nil
}

func GetBalance(projectId string) (*models.Balance, error) {
	rows, err := db.Query(`
	SELECT members.id, members.name, COALESCE(v_paid_balance.amount, 0) paid, COALESCE(v_spent_balance.amount, 0) spent 
		FROM members
  			LEFT JOIN v_paid_balance ON members.id = v_paid_balance.member_id
  			LEFT JOIN v_spent_balance ON members.id = v_spent_balance.member_id
			WHERE members.project_id = $1;`, projectId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get spent balance")
	}

	var balance []*models.MemberBalance
	for rows.Next() {
		var memberBalance models.MemberBalance
		err = rows.Scan(&memberBalance.Id, &memberBalance.Name, &memberBalance.Paid, &memberBalance.Spent)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		balance = append(balance, &memberBalance)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get spent balance")
	}
	return &models.Balance{Members: balance}, nil
}
