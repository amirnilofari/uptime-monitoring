package handlers

import (
	"net/http"

	"github.com/amirnilofari/uptime-monitoring-backend/db"
	"github.com/amirnilofari/uptime-monitoring-backend/models"
	"github.com/labstack/echo/v4"
)

func GetStatus(c echo.Context) error {
	userID := c.Get("user_id").(int)

	query := `
		SELECT us.id, us.url_id, us.status_code, us.response_time, us.checked_at, u.url 
		FROM url_status us JOIN urls u ON us.url_id = u.id
		WHERE u.user_id = $1
		ORDER BY us.checked_at DESC
	`

	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch status"})
	}
	defer rows.Close()

	type StatusResponse struct {
		models.URLStatus
		URL string `json:"url"`
	}

	var statuses []StatusResponse
	for rows.Next() {
		var status StatusResponse
		err := rows.Scan(&status.ID, &status.URLID, &status.StatusCode, &status.ResponseTime, &status.CheckedAt, &status.URL)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to parse status"})
		}
		statuses = append(statuses, status)

	}

	return c.JSON(http.StatusOK, echo.Map{"statuses": statuses})
}
