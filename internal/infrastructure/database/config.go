package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

const (
	MIGRATION_FILES_PATH = "./internal/infrastructure/database/sql"
)

func InitDatabase() *gorm.DB {
	username := os.Getenv("DB_USERNAME")
	if username == "" {
		log.Panic("Please set the DB_USERNAME in .env!")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Panic("Please set the DB_PASSWORD in .env!")
	}

	databaseName := os.Getenv("DB_DATABASE")
	if databaseName == "" {
		log.Panic("Please set the DB_DATABASE in .env!")
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Panic("Please set the DB_HOST in .env!")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		log.Panic("Please set the DB_PORT in .env!")
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Africa/Harare", host, username, password, databaseName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("Failed to connect database: %v", err)
	}

	Db = db
	return db
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
