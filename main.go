package main

import (
	"database/sql"
	"elite-consumer-service/internal/config"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	cfg := config.LoadConfig()
	fmt.Println("Database Url:", cfg.DatabaseUrl)

	db, err := sql.Open("postgres", cfg.DatabaseUrl)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Database ping failed", err)
	}

	fmt.Println("Successfully connected to the database")
	defer db.Close()

}
