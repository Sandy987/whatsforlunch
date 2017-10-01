package domain

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file" // Required to migrate from files
)

// MigrateToLatest will fetch the DB connection and attempt to migrate to the latest configuration
func MigrateToLatest() error {
	driver, err := postgres.WithInstance(DB.DB, &postgres.Config{})

	if err != nil {
		return err
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	fmt.Println(dir)

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)

	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil {
		return err
	}

	return nil
}
