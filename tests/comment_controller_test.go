package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aleksanderpalamar/backend-blog/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

var Router *gin.Engine

func TestCreateComment(t *testing.T) {
	token := "Bearer " + generateTestToken(1)

	user := models.User{Username: "testuser", Password: "password"}
	TestDB.Create(&user)

	post := models.Post{Title: "Test Post", Content: []string{"Test content"}}
	TestDB.Create(&post)

	w := httptest.NewRecorder()
	body := fmt.Sprintf(`{"post_id":%d,"content":"This is a comment"}`, post.ID)
	req, _ := http.NewRequest("POST", "/comments", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "This is a comment", response["data"].(map[string]interface{})["content"])
}

func TestFindComments(t *testing.T) {
	user := models.User{Username: "testuser", Password: "password"}
	TestDB.Create(&user)

	post := models.Post{Title: "Test Post", Content: []string{"Test content"}}
	TestDB.Create(&post)

	comment := models.Comment{PostID: post.ID, UserID: user.ID, Content: "Test comment"}
	TestDB.Create(&comment)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/posts/%d/comments", post.ID), nil)

	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	comments := response["data"].([]interface{})
	assert.Len(t, comments, 1)
	assert.Equal(t, "Test comment", comments[0].(map[string]interface{})["content"])
}

func generateTestToken(userID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})
	tokenString, _ := token.SignedString([]byte("secret"))
	return tokenString
}
