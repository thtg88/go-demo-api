package database

import (
	"fmt"
	"goDemoApi/app/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Instance Returns the existing GORM DB instance
func Instance() *gorm.DB {
	return db
}

// Connect to a Postgres database.
// The function fetches environment variables to compose the DSN string to connect
// And assigns the DB instance to the module db var
func Connect() *gorm.DB {
	var err error

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s %s TimeZone=UTC",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to the database")
	}

	return db
}

// AutoMigrate migrates the models specified
func AutoMigrate() {
	db.AutoMigrate(&models.User{})
}
