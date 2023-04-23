PRAGMA user_version = 1;

CREATE TABLE IF NOT EXISTS projects (
		id text NOT NULL PRIMARY KEY,
		name text NOT NULL,
		created_at text NOT NULL
    );

CREATE TABLE IF NOT EXISTS members (
		id integer NOT NULL PRIMARY KEY,
		name text NOT NULL,
		project_id text NOT NULL,
		FOREIGN KEY (project_id) REFERENCES projects(id)
	);

CREATE TABLE IF NOT EXISTS expenses (
		id integer NOT NULL PRIMARY KEY,
		project_id text,
		title text NOT NULL,
		amount real NOT NULL,
		paid_by integer,
		is_transfer bool NOT NULL,
		FOREIGN KEY (project_id) REFERENCES projects(id),
		FOREIGN KEY (paid_by) REFERENCES members(id)
	);

CREATE TABLE IF NOT EXISTS spent_by (
		expense_id integer,
		member_id integer,
		FOREIGN KEY (expense_id) REFERENCES expenses(id),
		FOREIGN KEY (member_id) REFERENCES members(id)
	);

CREATE VIEW IF NOT EXISTS v_parts AS
		SELECT expense_id, project_id, title, amount / COUNT(*) amount
			FROM expenses
			JOIN spent_by ON expenses.id = spent_by.expense_id
			GROUP BY expenses.id;

CREATE VIEW IF NOT EXISTS v_spent_balance AS
		SELECT spent_by.member_id, SUM(amount) amount
			FROM v_parts
			JOIN spent_by ON spent_by.expense_id = v_parts.expense_id
			GROUP BY spent_by.member_id;

CREATE VIEW IF NOT EXISTS v_paid_balance AS
		SELECT paid_by AS member_id, project_id, SUM(amount) amount 
            FROM expenses 
            GROUP BY paid_by;