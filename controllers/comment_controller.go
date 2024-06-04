package controllers

import (
	"net/http"

	"github.com/aleksanderpalamar/backend-blog/config"
	"github.com/aleksanderpalamar/backend-blog/models"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("user")
	input.UserID = user.(models.User).ID

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}
func FindCommentsByPost(c *gin.Context) {
	var comments []models.Comment
	if err := config.DB.Where("post_id = ?", c.Param("id")).Find(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}
