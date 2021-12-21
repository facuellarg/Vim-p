package main

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	//"github.com/golang-migrate/migrate/v4"
	//"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	//_ "github.com/lib/pq"
)

//migrationInstance get the instance for start migration
func migrationInstance(db *sql.DB) (*migrate.Migrate, error) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	return m, err
}

func MigrateUp(db *sql.DB) error {
	m, err := migrationInstance(db)
	if err != nil {
		return err
	}
	return m.Up()
}
func MigrateDown(db *sql.DB) error {
	m, err := migrationInstance(db)
	if err != nil {
		return err
	}
	return m.Down()
}

//MigrateRestart restart migrations
func MigrateRestart(db *sql.DB) error {
	err := MigrateDown(db)
	if err != nil {
		return err
	}
	return MigrateUp(db)
}
