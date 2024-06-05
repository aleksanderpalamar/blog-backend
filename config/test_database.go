package config

import (
	"log"

	"github.com/aleksanderpalamar/backend-blog/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
}

func SetupTestDatabase() *gorm.DB {
	testDB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to test database: %v", err)
	}

	err = testDB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatalf("Error migrating test database: %v", err)
	}

	return testDB
}
