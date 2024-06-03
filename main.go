package main

import (
	"log"
	"os"

	"github.com/aleksanderpalamar/backend-blog/config"
	"github.com/aleksanderpalamar/backend-blog/middleware"
	"github.com/aleksanderpalamar/backend-blog/models"
	"github.com/aleksanderpalamar/backend-blog/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Load variables from .env file
	config.LoadEnvVariables()
	// Connect to database
	config.ConnectDatabase()
	// Migrate models
	config.DB.AutoMigrate(&models.Post{})
	// Config routes
	r = routes.SetupRouter()
	// Config CORS
	r.Use(middleware.CORSMiddleware())
	// Define PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Run the application
	log.Fatal(r.Run(":" + port))
}
