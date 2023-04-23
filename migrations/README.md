# SQLite Migrations

To add a database migration, create a `.sql` file. 
- The file MUST have the name of the "from" version.
- The content of the file MUST be a SQL TRANSACTION
- Within this transaction, update the version using `PRAGMA user_version = <NEW_VERSION>;`