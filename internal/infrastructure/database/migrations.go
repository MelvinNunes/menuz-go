package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	db, _ := sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, host, port, databaseName))
	driver, _ := postgres.WithInstance(db, &postgres.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/infrastructure/database/migrations",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("No change made by migration scripts")
		} else {
			log.Fatal(err)
		}
	}
}
