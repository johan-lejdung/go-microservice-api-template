package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Need to import this package to access database driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
)

// CreateDatabase will create and return a database
func CreateDatabase() (*sql.DB, error) {
	// Load env variables
	serverName := os.Getenv("DB_SERVERNAME")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", user+":"+password+"@tcp("+serverName+")/"+dbName+"?&charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true")
	if err != nil {
		return nil, err
	}

	if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "test" {
		err := migrateDatabase(db)
		if err != nil {
			return db, err
		}
	}

	return db, nil
}

func migrateDatabase(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/db/migrations", dir),
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	migration.Log = &MigrationLogger{}

	migration.Log.Printf("Applying database migrations")
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	version, _, err := migration.Version()
	if err != nil {
		return err
	}

	migration.Log.Printf("Active database version: %d", version)

	return nil
}
