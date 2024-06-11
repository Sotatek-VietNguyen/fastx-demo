package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"user-service/config"
	"user-service/src/pkg/database"
	"user-service/src/routes"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Set up Gin router
	r := gin.Default()

	// Initialize routes
	routes.InitializeRoutes(r, db)

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
