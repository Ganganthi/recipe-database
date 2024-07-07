package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	dbModels "main/internal/database/models"
)

var DbConn *gorm.DB = nil

func Connect() error {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf(
		"host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable",
		user,
		password,
		dbName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return err
	}

	// Migrate the database after connecting
	err = db.AutoMigrate(&dbModels.User{}, &dbModels.Ingredient{})
	if err != nil {
		log.Println(err)
		return err
	}

	DbConn = db
	return nil
}
