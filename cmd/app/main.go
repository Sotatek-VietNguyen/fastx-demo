package main

import (
	"fastx-api/config"
	"fastx-api/src/pkg/database"
	"fastx-api/src/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Run seeders
	//err = seeders.SeedUsers(db)
	//if err != nil {
	//	log.Fatalf("Could not seed users: %v", err)
	//}

	// Set up Gin router
	r := gin.Default()

	// Initialize routes
	routes.InitializeRoutes(r, db)

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
