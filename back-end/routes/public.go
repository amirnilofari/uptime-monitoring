package routes

import (
	"github.com/amirnilofari/uptime-monitoring-backend/handlers"
	"github.com/labstack/echo/v4"
)

func PublicRoutes(router *echo.Echo) {
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
}
