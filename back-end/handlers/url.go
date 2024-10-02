package handlers

import (
	"net/http"
	"time"

	"github.com/amirnilofari/uptime-monitoring-backend/db"
	"github.com/amirnilofari/uptime-monitoring-backend/models"
	"github.com/labstack/echo/v4"
)

func AddURL(c echo.Context) error {
	userID := c.Get("user_id").(int)

	type Request struct {
		URL           string `json:"url"`
		CheckInterval int    `json:"check_interval"`
	}

	req := new(Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	url := models.URL{
		UserID:        userID,
		URL:           req.URL,
		CheckInterval: req.CheckInterval,
		CreatedAt:     time.Now(),
	}

	query := "INSERT INTO urls (user_id, url, check_interval, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	err := db.DB.QueryRow(query, url.UserID, url.URL, url.CheckInterval, url.CreatedAt).Scan(&url.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to add URL"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"url": url})
}

func GetURLs(c echo.Context) error {
	userID := c.Get("user_id").(int)

	query := "SELECT id, url, check_interval, created_at FROM urls WHERE user_id=$1"
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch urls"})
	}
	defer rows.Close()

	var urls []models.URL
	for rows.Next() {
		var url models.URL
		err := rows.Scan(&url.ID, &url.URL, &url.CheckInterval, &url.CreatedAt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to parse urls"})
		}
		urls = append(urls, url)
	}

	return c.JSON(http.StatusOK, echo.Map{"urls": urls})
}

func DeleteURL(c echo.Context) error {
	userID := c.Get("user_id").(int)
	urlID := c.Param("id")

	query := "DELETE FROM urls WHERE id=$1 AND user_id=$2"
	res, err := db.DB.Exec(query, urlID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete url" + err.Error()})
	}

	count, err := res.RowsAffected()
	if err != nil || count == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "url not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "url deleted successfully!"})
}
