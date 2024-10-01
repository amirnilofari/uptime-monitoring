package monitor

import (
	"net/http"
	"time"

	"github.com/amirnilofari/uptime-monitoring-backend/db"
	"github.com/amirnilofari/uptime-monitoring-backend/models"
)

func StartMonitoring() {
	for {
		scheduleChecks()
		time.Sleep(1 * time.Minute)
	}
}

var lastChecked = make(map[int]time.Time)

func scheduleChecks() {
	query := "SELECT id, url, check_interval FROM urls"
	rows, err := db.DB.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	now := time.Now()
	for rows.Next() {
		var url models.URL
		err := rows.Scan(&url.ID, &url.URL, &url.CheckInterval)
		if err != nil {
			continue
		}

		last, exists := lastChecked[url.ID]
		if !exists || now.Sub(last).Minutes() >= float64(url.CheckInterval) {
			go checkURL(url)
			lastChecked[url.ID] = now
		}
	}

}

func checkURL(url models.URL) {
	start := time.Now()
	resp, err := http.Get(url.URL)
	responseTime := int(time.Since(start).Milliseconds())

	statusCode := 0

	if err == nil {
		statusCode = resp.StatusCode
		resp.Body.Close()
	}

	query := "INSERT INTO url_status (url_id, status_code, response_time, checked_at) VALUES ($1, $2, $3, $4)"
	db.DB.Exec(query, url.ID, statusCode, responseTime, time.Now())
}
