package config

import (
	"github.com/aleksanderpalamar/backend-blog/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate models
	database.AutoMigrate(&models.Post{}, &models.Comment{}, &models.User{})

	DB = database

	{
		/*
				var err error
			dsn := fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
				os.Getenv("DB_HOST"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_NAME"),
			)
			DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				panic("Failed to connect to database!")
			}
		*/
	}
}
