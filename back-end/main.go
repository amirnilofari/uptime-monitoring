package main

import (
	"log"
	"os"

	"github.com/amirnilofari/uptime-monitoring-backend/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	db_url := os.Getenv("DB_URL")
	err = db.InitDB(db_url)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
}
