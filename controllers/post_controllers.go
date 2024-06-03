package controllers

import (
	"net/http"

	"github.com/aleksanderpalamar/backend-blog/config"
	"github.com/aleksanderpalamar/backend-blog/models"
	"github.com/gin-gonic/gin"
)

func FindAllPosts(context *gin.Context) {
	var posts []models.Post
	config.DB.Find(&posts)
	context.JSON(http.StatusOK, gin.H{"data": posts})
}

func FindPost(context *gin.Context) {
	var post models.Post
	if err := config.DB.First(&post, context.Param("id")).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Post not found!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": post})
}

func CreatePost(context *gin.Context) {
	var input models.Post
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	context.JSON(http.StatusOK, gin.H{"data": input})
}

func UpdatePost(context *gin.Context) {
	var post models.Post
	if err := config.DB.First(&post, context.Param("id")).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Post not found!"})
		return
	}

	var input models.Post
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&post).Updates(input)
	context.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(context *gin.Context) {
	var post models.Post
	if err := config.DB.First(&post, context.Param("id")).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Post not found!"})
		return
	}

	config.DB.Delete(&post)
	context.JSON(http.StatusOK, gin.H{"data": true})
}
