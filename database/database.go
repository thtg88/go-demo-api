package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Instance() *gorm.DB {
	return db
}

func Connect() *gorm.DB {
	var err error

	dsn := fmt.Sprintf(
		"user=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to the database")
	}

	return db
}
