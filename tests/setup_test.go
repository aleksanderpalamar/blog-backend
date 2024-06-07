package tests

import (
	"os"
	"testing"

	"github.com/aleksanderpalamar/backend-blog/config"
	"github.com/aleksanderpalamar/backend-blog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func SetupTestDatabase() {
	config.ConnectTestDatabase()
	if config.TestDB == nil {
		panic("Failed to connect to test database")
	}
	config.TestDB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestMain(m *testing.M) {
	SetupTestDatabase()
	code := m.Run()
	os.Exit(code)
}
