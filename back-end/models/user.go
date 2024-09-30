package models

import "time"

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"last_name"`
	LastName  string    `json:"first_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password_hash"`
	CreatedAt time.Time `json:"created_at"`
}
