package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

var db *sql.DB

func Setup() (func() error, error) {
	err := os.Mkdir("data", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return nil, errors.Wrap(err, "failed to create folder")
	}

	db, err = sql.Open("sqlite3", "data/data.db")
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database")
	}

	err = runAllMigrations()
	if err != nil {
		return nil, errors.Wrap(err, "failed to migrate database")
	}

	return db.Close, nil
}

// Execute all migrations in order, by checking if the migration script exists for each given version.
// For example, if the current user_version is 5, checks if 5.sql exists, and if so, execute it.
// Cyclic migrations are also detected and throw an error (happens if a migration does not update user_version)
func runAllMigrations() error {
	usedFilenames := make(map[string]bool)
	for {
		userVersion, err := getUserVersion()
		if err != nil {
			return errors.Wrap(err, "failed to get migration filename")
		}
		log.Printf("INFO: Current database user_version is %v", userVersion)
		filename := fmt.Sprintf("migrations/%v.sql", userVersion)

		if usedFilenames[filename] {
			return fmt.Errorf("cyclic migration detected, does %s upgrade to a new user_version ?", filename)
		}
		usedFilenames[filename] = true

		hasMigration := fileExists(filename)
		if !hasMigration {
			fmt.Printf("INFO: Database migrations completed: %s does not exists\n", filename)
			break
		}

		log.Printf("INFO: Executing migration with file %s...", filename)
		err = runMigrationScript(filename)
		if err != nil {
			return errors.Wrapf(err, "failed to run migration %s", filename)
		}
		log.Printf("INFO: Migration successful !")
	}
	return nil
}

func getUserVersion() (int, error) {
	row := db.QueryRow("PRAGMA user_version")
	var userVersion int
	err := row.Scan(&userVersion)
	if err != nil {
		return 0, errors.Wrap(err, "failed to read user_version")
	}
	return userVersion, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func runMigrationScript(filename string) error {
	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "failed to read migration file")
	}
	contentStr := string(contentBytes)

	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		return errors.Wrap(err, "failed to begin tx")
	}

	_, err = tx.Exec(contentStr)
	if err != nil {
		return errors.Wrap(err, "failed to execute migration")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit migration")
	}
	return nil
}
