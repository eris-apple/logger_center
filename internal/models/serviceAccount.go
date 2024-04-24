package models

import "time"

type ServiceAccount struct {
	ID        string    `json:"id"`
	ProjectID string    `json:"project_id"`
	IsActive  bool      `json:"is_active"`
	Secret    string    `json:"secret"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
