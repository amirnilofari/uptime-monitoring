package main

import (
	"log"
	"os"

	"github.com/amirnilofari/uptime-monitoring-backend/db"
	"github.com/amirnilofari/uptime-monitoring-backend/monitor"
	"github.com/amirnilofari/uptime-monitoring-backend/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
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

	e := echo.New()
	e.Use(middleware.CORS())
	//e.Use(middleware.CORS())

	routes.PublicRoutes(e)
	routes.ProtectedRoutes(e)

	go monitor.StartMonitoring()

	e.Logger.Fatal(e.Start(":8080"))

}
