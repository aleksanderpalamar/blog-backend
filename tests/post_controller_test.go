package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aleksanderpalamar/backend-blog/controllers"
	"github.com/aleksanderpalamar/backend-blog/middleware"
	"github.com/aleksanderpalamar/backend-blog/models"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	router := SetupTestEnvironment()
	router.POST("/posts", middleware.Auth(), controllers.CreatePost)

	token := "Bearer " + generateTestToken(1)

	w := httptest.NewRecorder()
	body := `{"title":"New Post","content":["This is a new post"]}`
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "New Post", response["data"].(map[string]interface{})["title"])
}

func TestFindPosts(t *testing.T) {
	router := SetupTestEnvironment()

	router.GET("/posts", controllers.FindAllPosts)

	post := models.Post{Title: "Test Post", Content: []string{"Test content"}}
	TestDB.Create(&post)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	posts := response["data"].([]interface{})
	assert.Len(t, posts, 1)
	assert.Equal(t, "Test Post", posts[0].(map[string]interface{})["title"])
}
