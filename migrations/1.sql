PRAGMA user_version = 2;

ALTER TABLE expenses 
    ADD COLUMN updated_at2 text NOT NULL DEFAULT "Sun Apr 23 12:15:55 CEST 2023";