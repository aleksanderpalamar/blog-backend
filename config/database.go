package config

import (
	"fmt"
	"log"
	"os"

	"github.com/aleksanderpalamar/backend-blog/models"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db

	// Migrate models
	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatalf("Failed to migrate database models: %v", err)
	}

	log.Println("Database connection established successfully")
}
