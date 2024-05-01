package models

import (
	"time"
)

type Log struct {
	ID        string    `json:"id"`
	ChainID   string    `json:"chain_id"`
	ProjectID string    `json:"project_id"`
	Content   string    `json:"content"`
	Level     string    `json:"level"`
	Timestamp time.Time `json:"timestamp"`
	CreatedAt time.Time `json:"created_at"`
}
