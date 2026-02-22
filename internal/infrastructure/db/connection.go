package db

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func getDBEnvVars() (user, pass, host, port, name string, err error) {
	user = os.Getenv("DB_USER")
	pass = os.Getenv("DB_PASSWORD")
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	name = os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		err = fmt.Errorf("missing one or more env variables: DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME")
	}

	return
}

func RunMigrations() error {
	dsn, err := buildDSN()
	if err != nil {
		return err
	}

	m, err := migrate.New(
		"file:///app/migrations",
		dsn,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func buildDSN() (string, error) {
	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		return "", fmt.Errorf("missing env variable: DB_DRIVER")
	}

	user, pass, host, port, name, err := getDBEnvVars()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s", driver, user, pass, host, port, name), nil
}

func NewConnection() (*sqlx.DB, error) {
	user, pass, host, port, name, err := getDBEnvVars()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
