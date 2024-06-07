package routes

import (
	"github.com/aleksanderpalamar/backend-blog/controllers"
	"github.com/aleksanderpalamar/backend-blog/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.InjectDB(db))
	router.Use(middleware.CORSMiddleware())

	// Public Routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Protected Routes
	protected := router.Group("/")
	protected.Use(middleware.Auth())

	protected.GET("/posts", controllers.FindAllPosts)
	protected.GET("/posts/:id", controllers.FindPost)
	protected.POST("/posts", controllers.CreatePost)
	protected.PUT("/posts/:id", controllers.UpdatePost)
	protected.DELETE("/posts/:id", controllers.DeletePost)

	protected.GET("/posts/:id/comments", controllers.FindCommentsByPost)
	protected.POST("/posts/:id/comments", controllers.CreateComment)

	return router
}
