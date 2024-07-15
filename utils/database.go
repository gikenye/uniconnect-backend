package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"uniconnect/models"
)

var DB *gorm.DB

func InitialiseDB() {
	var err error
	log.Print("Initialising Database...")

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Could not find Database Url")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	log.Print("Successfully connected database!")

	setupModels(
		&models.User{},
		&models.Business{},
		&models.Likes{},
		&models.Comment{},
	)
}

func setupModels(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
}
