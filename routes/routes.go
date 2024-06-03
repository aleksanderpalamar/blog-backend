package routes

import (
	"github.com/aleksanderpalamar/backend-blog/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/posts", controllers.FindAllPosts)
	router.GET("/posts/:id", controllers.FindPost)
	router.POST("/posts", controllers.CreatePost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	return router
}
