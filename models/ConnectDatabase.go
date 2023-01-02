package models

import (
	"api/constants"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		constants.DB_HOST,
		constants.DB_USER,
		constants.DB_PASSWORD,
		constants.DB_NAME,
		constants.DB_PORT,
		constants.DB_SSLMODE,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
