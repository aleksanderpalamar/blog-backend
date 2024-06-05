package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aleksanderpalamar/backend-blog/controllers"
	"github.com/aleksanderpalamar/backend-blog/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestRegisterUser(t *testing.T) {
	router := SetupTestEnvironment()
	router.POST("/register", controllers.Register)

	w := httptest.NewRecorder()
	body := `{"username":"testuser","password":"password123"}`
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "testuser", response["data"].(map[string]interface{})["username"])
}

func TestLoginUser(t *testing.T) {
	router := SetupTestEnvironment()
	router.POST("/login", controllers.Login)

	// Create a test user
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := models.User{Username: "testuser", Password: string(hashedPassword)}
	TestDB.Create(&user)

	w := httptest.NewRecorder()
	body := `{"username":"testuser","password":"password123"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.NotEmpty(t, response["token"])
}
