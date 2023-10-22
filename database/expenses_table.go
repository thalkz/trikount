package database

import (
	"time"

	"github.com/pkg/errors"
	"github.com/thalkz/trikount/models"
)

func AddExpense(projectId string, title string, amount float64, paidBy int, spendBy []int, isTransfer bool, createdAt time.Time) error {
	row := db.QueryRow(`INSERT INTO expenses (title, amount, project_id, paid_by, is_transfer, created_at)
		values($1, $2, $3, $4, $5, $6) RETURNING id`, title, amount, projectId, paidBy, isTransfer, createdAt.Format(time.DateTime))
	var expenseId int
	err := row.Scan(&expenseId)
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

func EditExpense(projectId string, expenseId int, title string, amount float64, paidBy int, spentBy []int, isTransfer bool, createdAt time.Time) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin tx")
	}
	defer tx.Rollback()

	_, err = tx.Exec(`UPDATE expenses 
		SET title = ?, amount = ?, paid_by = ?, is_transfer = ?, created_at = ? 
		WHERE id = ? AND project_id = ?`,
		title, amount, paidBy, isTransfer, createdAt.Format(time.DateTime), expenseId, projectId)
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

func DeleteExpense(projectId string, expenseId int) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin tx")
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM expenses WHERE id = ? AND project_id = ?", expenseId, projectId)
	if err != nil {
		return errors.Wrap(err, "failed to delete from expenses")
	}

	_, err = tx.Exec("DELETE FROM spent_by WHERE expense_id = ?", expenseId)
	if err != nil {
		return errors.Wrap(err, "failed to delete from spent_by")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit tx")
	}

	return nil
}

func GetExpenses(projectId string) ([]*models.Expense, error) {
	rows, err := db.Query(`SELECT expenses.id, expenses.title, expenses.amount, members.id, members.name, expenses.is_transfer, expenses.created_at
		FROM expenses 
		JOIN members ON expenses.paid_by = members.id
		WHERE expenses.project_id = $1
		ORDER BY expenses.created_at desc`, projectId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get expenses")
	}

	var expenses []*models.Expense
	for rows.Next() {
		var expense models.Expense
		var createdAtStr string
		err := rows.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.PaidBy.Id, &expense.PaidBy.Name, &expense.IsTransfer, &createdAtStr)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan expense row")
		}
		expense.CreatedAt, err = time.Parse(time.DateTime, createdAtStr)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse time")
		}
		expenses = append(expenses, &expense)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan expenses")
	}

	return expenses, nil
}

func GetExpenseWithParts(projectId string, memberId int) ([]*models.ExpenseWithPart, error) {
	rows, err := db.Query(`SELECT expenses.id, expenses.title, expenses.amount, members.id, members.name, expenses.is_transfer, expenses.created_at, iif(user_spent.member_id = ? AND v_parts.is_transfer = False, v_parts.amount, 0) as part
			FROM expenses
			JOIN members ON expenses.paid_by = members.id
			JOIN v_parts ON v_parts.expense_id = expenses.id		
			LEFT JOIN (SELECT * FROM spent_by WHERE member_id = ?) as user_spent ON expenses.id = user_spent.expense_id
		WHERE expenses.project_id = ?
		ORDER BY expenses.created_at desc`, memberId, memberId, projectId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get expense with parts")
	}

	var expenseWithParts []*models.ExpenseWithPart
	for rows.Next() {
		var expense models.ExpenseWithPart
		var createdAtStr string
		err := rows.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.PaidBy.Id, &expense.PaidBy.Name, &expense.IsTransfer, &createdAtStr, &expense.Part)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan expense row")
		}
		expense.CreatedAt, err = time.Parse(time.DateTime, createdAtStr)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse time")
		}
		expenseWithParts = append(expenseWithParts, &expense)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan expense with parts")
	}

	return expenseWithParts, nil
}

func GetTotalSpent(projectId string) (float64, error) {
	row := db.QueryRow(`SELECT COALESCE(SUM(amount), 0.0) AS total FROM expenses WHERE project_id = $1 AND is_transfer = FALSE`, projectId)
	var total float64
	err := row.Scan(&total)
	if err != nil {
		return 0.0, errors.Wrap(err, "failed to scan total expense")
	}
	return total, nil
}

func GetExpense(projectId string, id int) (*models.Expense, error) {
	row := db.QueryRow(`SELECT expenses.id, expenses.title, expenses.amount, members.id, members.name, expenses.is_transfer, expenses.created_at
		FROM expenses 
		JOIN members ON expenses.paid_by = members.id
		WHERE expenses.id = $1 AND expenses.project_id = $2`, id, projectId)

	var expense models.Expense
	var createdAtStr string
	err := row.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.PaidBy.Id, &expense.PaidBy.Name, &expense.IsTransfer, &createdAtStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan expense")
	}
	expense.CreatedAt, err = time.Parse(time.DateTime, createdAtStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse time")
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
		expense.SpentBy = append(expense.SpentBy, &member)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get spent_by members")
	}

	return &expense, nil
}

func GetBalance(projectId string) (*models.Balance, error) {
	rows, err := db.Query(`
	SELECT members.id, members.name, COALESCE(v_paid_balance.amount, 0) paid, COALESCE(v_spent_balance.amount, 0) spent, COALESCE(v_spent_balance.no_transfer_amount, 0) no_transfer_spent 
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
		err = rows.Scan(&memberBalance.Id, &memberBalance.Name, &memberBalance.Paid, &memberBalance.Spent, &memberBalance.NoTransferSpent)
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
