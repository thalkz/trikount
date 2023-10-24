-- WARNING: THIS STEP IS OVERRIDEN BY 6.sql

PRAGMA user_version = 6;

UPDATE expenses
    SET created_at = created_at || ':00'
