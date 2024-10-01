package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/amirnilofari/uptime-monitoring-backend/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "No authorization header provided!"})
		}

		tokenString = strings.TrimSpace(strings.Replace(tokenString, "Bearer", "", 1))

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return utils.JwtSecretKey, nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invaild token"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, ok := claims["user_id"].(float64)
			if ok {
				c.Set("user_id", int(userID))
			} else {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid user ID in token"})
			}

		} else {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
		}

		return next(c)
	}
}
