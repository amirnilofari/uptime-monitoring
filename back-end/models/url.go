package models

import "time"

type URL struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	URL           string    `json:"url"`
	CheckInterval int       `json:"check_interval"`
	CreatedAt     time.Time `json:"created_at"`
}
