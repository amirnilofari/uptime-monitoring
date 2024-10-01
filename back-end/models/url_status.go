package models

import "time"

type URLStatus struct {
	ID           int       `json:"id"`
	URLID        int       `json:"url_id"`
	StatusCode   int       `json:"status_code"`
	ResponseTime int       `json:"response_time"`
	CheckedAt    time.Time `json:"checked_at"`
}
