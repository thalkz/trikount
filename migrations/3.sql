PRAGMA user_version = 4;

DROP VIEW v_parts;

CREATE VIEW IF NOT EXISTS v_parts AS
		SELECT expense_id, project_id, title, amount / COUNT(*) amount, is_transfer
			FROM expenses
			JOIN spent_by ON expenses.id = spent_by.expense_id
			GROUP BY expenses.id;

DROP VIEW v_spent_balance;

CREATE VIEW IF NOT EXISTS v_spent_balance AS
		SELECT spent_by.member_id, SUM(amount) amount, SUM(iif(NOT is_transfer, amount, 0)) no_transfer_amount
			FROM v_parts
			JOIN spent_by ON spent_by.expense_id = v_parts.expense_id
			GROUP BY spent_by.member_id;