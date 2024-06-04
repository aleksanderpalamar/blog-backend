package routes

import (
	"github.com/aleksanderpalamar/backend-blog/controllers"
	"github.com/aleksanderpalamar/backend-blog/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	router.GET("/posts", controllers.FindAllPosts)
	router.GET("/posts/:id", controllers.FindPost)
	router.POST("/posts", middleware.Auth(), controllers.CreatePost)
	router.PUT("/posts/:id", middleware.Auth(), controllers.UpdatePost)
	router.DELETE("/posts/:id", middleware.Auth(), controllers.DeletePost)

	router.GET("/posts/:id/comments", controllers.FindCommentsByPost)
	router.POST("/posts/:id/comments", middleware.Auth(), controllers.CreateComment)

	return router
}
