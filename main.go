package main

import (
	"log"
	"os"

	"github.com/aleksanderpalamar/backend-blog/config"
	"github.com/aleksanderpalamar/backend-blog/middleware"
	"github.com/aleksanderpalamar/backend-blog/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	mode := os.Getenv("MODE")
	if mode == "production" {
		mode = "release"
	} else {
		mode = "debug"
		log.Println("Running in debug mode")
	}
	gin.SetMode(mode)

	r := gin.Default()
	// Cors middleware
	middleware.CORSMiddleware()
	// Load variables from .env file
	config.LoadEnvVariables()
	// Connect to database
	config.ConnectDatabase()
	// Config routes
	r = routes.SetupRouter(
		config.DB,
	)
	// Define PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Run the application
	log.Fatal(r.Run(":" + port))
}
