PRAGMA user_version = 7;

-- 1) Drop the previous created_at column (from 4.sql and 5.sql)
ALTER TABLE
    expenses DROP COLUMN created_at;

-- 2) Re-reate the column with the correct default
ALTER TABLE
    expenses
ADD
    COLUMN created_at text NOT NULL DEFAULT "1970-01-01 00:00:00";

-- 2) Re-update the column with the correct time-format
WITH expenses_date AS (
    SELECT
        id,
        trim(substr(updated_at, 9, 3)) AS day,
        trim(substr(updated_at, 12, 5)) AS time,
        CASE
            trim(substr(updated_at, 4, 4))
            WHEN 'Jan' THEN '01'
            WHEN 'Feb' THEN '02'
            WHEN 'Mar' THEN '03'
            WHEN 'Apr' THEN '04'
            WHEN 'Mai' THEN '05'
            WHEN 'Jun' THEN '06'
            WHEN 'Jul' THEN '07'
            WHEN 'Aug' THEN '08'
            WHEN 'Sep' THEN '09'
            WHEN 'Oct' THEN '10'
            WHEN 'Nov' THEN '11'
            WHEN 'Dev' THEN '12'
        END AS month,
        trim(substr(updated_at, length(updated_at) -4, 8)) AS year
    FROM
        expenses
)
UPDATE
    expenses
SET
    created_at = (
        SELECT
            year || '-' || month || '-' || substr('0' || day, -2, 2) || ' ' || time || ":00"
        FROM
            expenses_date
        WHERE
            expenses_date.id = expenses.id
    )