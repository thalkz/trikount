PRAGMA user_version = 3;

ALTER TABLE expenses 
    ADD COLUMN updated_at text NOT NULL DEFAULT "Sun Apr 23 12:15:55 CEST 2023";

ALTER TABLE expenses 
    DROP COLUMN updated_at2;