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
	var dialector gorm.Dialector

	dialector, err = GetDialector(os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic(err)
	}

	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("Error connecting to the database")
	}

	return db
}

// GetDialector returns a GORM Dialector from a given connection type string.
func GetDialector(connection string) (gorm.Dialector, error) {
	if connection != "pgsql" {
		return nil, fmt.Errorf("database: unrecognised connection type %s", connection)
	}

	dsn, _ := GetDsn(connection)

	return postgres.Open(dsn), nil
}

// GetDsn returns a DSN from a given database connection type string
func GetDsn(connection string) (string, error) {
	if connection != "pgsql" {
		return "", fmt.Errorf("database: unrecognised connection type %s", connection)
	}

	password := os.Getenv("DB_PASSWORD")
	sslMode := os.Getenv("DB_SSL_MODE")
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s",
		os.Getenv("DB_HOSTNAME"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
	)
	if password != "" {
		dsn = dsn + fmt.Sprintf(" password=%s", password)
	}
	if sslMode != "" {
		dsn = dsn + fmt.Sprintf(" sslmode=%s", sslMode)
	}
	dsn = dsn + " TimeZone=UTC"

	return dsn, nil
}

// AutoMigrate migrates the models specified
func AutoMigrate() {
	db.AutoMigrate(&models.User{})
}
