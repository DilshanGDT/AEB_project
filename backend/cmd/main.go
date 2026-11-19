package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"aeb-project/backend/configs"
	"aeb-project/backend/internal/database"
	"aeb-project/backend/internal/routes"
)

func main() {

	// Load configuration
	cfg := configs.LoadConfig()

	// Connect to MySQL
	db := database.ConnectDB(cfg)
	defer db.Close()

	// Initialize Gin
	router := gin.Default()

	// Configure routes
	routes.SetupRoutes(router, db)

	// Start HTTP server
	log.Println("Successfully connected to MySQL database!")
	log.Println("AEB Bond Management Backend Started")
	log.Println("Server running at http://localhost:8080")

	err := router.Run(":8080")

	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
