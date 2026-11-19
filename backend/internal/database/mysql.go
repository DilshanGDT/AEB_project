package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"aeb-project/backend/configs"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(cfg *configs.Config) *sql.DB {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Verify database connection
	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatal("Unable to connect to MySQL:", err)
	}

	log.Println("Successfully connected to MySQL database!")

	return db
}
