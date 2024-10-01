package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecretKey []byte

func GenerateJWT(userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(JwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
