package main

import (
	"log"

	"aeb-project/backend/configs"
	"aeb-project/backend/internal/database"
)

func main() {

	// Load application configuration
	cfg := configs.LoadConfig()

	// Connect to MySQL
	db := database.ConnectDB(cfg)
	defer db.Close()

	log.Println("AEB Bond Management Backend Started")

	// Your application initialization will go here.
	_ = db

	var dbName string

	err := db.QueryRow("SELECT DATABASE()").Scan(&dbName)
	if err != nil {
		log.Fatal("Database verification failed:", err)
	}

	log.Println("Connected Database:", dbName)
}
