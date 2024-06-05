package tests

import (
	"os"
	"testing"

	"github.com/aleksanderpalamar/backend-blog/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func SetupTestEnvironment() *gin.Engine {
	TestDB = config.SetupTestDatabase()

	router := gin.Default()
	return router
}

func TeardownTestEnvironment() {
	db, _ := TestDB.DB()
	db.Close()
}

func TestMain(m *testing.M) {
	SetupTestEnvironment()
	gin.SetMode(gin.TestMode)

	code := m.Run()
	TeardownTestEnvironment()
	os.Exit(code)
}
