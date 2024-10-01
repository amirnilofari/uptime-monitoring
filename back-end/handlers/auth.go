package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/amirnilofari/uptime-monitoring-backend/db"
	"github.com/amirnilofari/uptime-monitoring-backend/models"
	"github.com/amirnilofari/uptime-monitoring-backend/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	type Request struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password_hash"`
	}

	req := new(Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
	}

	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	query := `INSERT INTO users (first_name,last_name, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = db.DB.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt).Scan(&user.ID)
	if err != nil {
		return c.JSON(http.StatusConflict, echo.Map{"error": "User already exists!"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "User registered successfully!"})
}

func Login(c echo.Context) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password_hash"`
	}

	req := new(Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input!"})
	}

	var user models.User
	query := "SELECT id, first_name, last_name, password_hash FROM users WHERE email=$1"
	err := db.DB.QueryRow(query, req.Email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Server error"})
	}

	match := CheckPasswordHash(req.Password, user.Password)
	if match {
		token, err := utils.GenerateJWT(user.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not generate token"})
		}
		return c.JSON(http.StatusOK, echo.Map{"token": token})
	} else {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid password"})
	}

	//return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println("error to compare:", err)
		return false
	} else {
		return true
	}
}
