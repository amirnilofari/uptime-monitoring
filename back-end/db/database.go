package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(dsn string) error {
	var err error

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("Failed to open a DB connection: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %w", err)
	}

	log.Println("Database connection established!")
	return nil
}
