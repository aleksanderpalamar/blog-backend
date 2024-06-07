package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectTestDatabase() {
	var err error
	TestDB, err = gorm.Open(sqlite.Open("test_blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
