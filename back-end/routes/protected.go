package routes

import (
	"github.com/amirnilofari/uptime-monitoring-backend/handlers"
	"github.com/amirnilofari/uptime-monitoring-backend/middlewares"
	"github.com/labstack/echo/v4"
)

func ProtectedRoutes(router *echo.Echo) {
	r := router.Group("/")
	r.Use(middlewares.JWTAuthMiddleware)
	r.POST("/urls", handlers.AddURL)
	r.GET("/urls", handlers.GetURLs)
	r.DELETE("/urls/:id", handlers.DeleteURL)
	r.GET("/status", handlers.GetStatus)
}
