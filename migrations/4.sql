-- WARNING: THIS STEP IS OVERRIDEN BY 6.sql

PRAGMA user_version = 5;

-- STEP 1) Add a "created_at" column, that uses the UnixDate format
ALTER TABLE
    expenses
ADD
    COLUMN created_at text NOT NULL DEFAULT "1970-01-01 00:00:00";

-- STEP 2) Create an expenses_date view that parses the "updated_at" column
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

-- STEP 3) Update the new "created_at" column, in UnixDate format
UPDATE
    expenses
SET
    created_at = (
        SELECT
            year || '-' || month || '-' || substr('0' || day, -2, 2) || ' ' || time
        FROM
            expenses_date
        WHERE
            expenses_date.id = expenses.id
    )