package models

import "time"

type Session struct {
	ID        string    `json:"id"`
	Token     string    `json:"token"`
	IsActive  bool      `json:"is_active"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
