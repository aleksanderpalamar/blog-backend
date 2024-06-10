package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	"github.com/aleksanderpalamar/backend-blog/config"
	"github.com/aleksanderpalamar/backend-blog/controllers"
	"github.com/aleksanderpalamar/backend-blog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestRouter() *gin.Engine {
	database, _ := gorm.Open(sqlite.Open("test_blog.db"), &gorm.Config{})
	config.DB = database
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	// Middleware to provide DB instance
	r.Use(func(c *gin.Context) {
		c.Set("db", config.DB)
		c.Next()
	})

	return r
}

func TestRegister(t *testing.T) {
	router := setupTestRouter()
	router.POST("/register", controllers.Register)

	payload := `{"username":"testuser", "password":"password123"}`
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()), &response)

	data := response["data"].(map[string]interface{})
	assert.Equal(t, "testuser", data["username"])

	// Try to register the same user again
	req, _ = http.NewRequest("POST", "/register", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, "Username already taken", response["error"])
}

func TestLogin(t *testing.T) {
	router := setupTestRouter()

	// Create a user for testing login
	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := models.User{Username: "testuser", Password: string(password)}
	config.DB.Create(&user)

	router.POST("/login", controllers.Login)

	payload := `{"username":"testuser", "password":"password123"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()), &response)

	token := response["token"].(string)
	assert.NotEmpty(t, token)
}

func TestLoginInvalidCredentials(t *testing.T) {
	router := setupTestRouter()

	// Create a user for testing login
	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := models.User{Username: "testuser", Password: string(password)}
	config.DB.Create(&user)

	router.POST("/login", controllers.Login)

	payload := `{"username":"testuser", "password":"wrongpassword"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]interface{}
	json.Unmarshal([]byte(w.Body.String()), &response)

	errorMessage := response["error"].(string)
	assert.Equal(t, "Invalid username or password", errorMessage)
}
