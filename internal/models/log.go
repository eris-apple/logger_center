package models

import (
	"time"
)

type Log struct {
	ID        string    `json:"id"`
	ChainID   string    `json:"chain_id"`
	ProjectID string    `json:"project_id"`
	Title     string    `json:"title"`
	Error     string    `json:"error"`
	Params    string    `json:"params"`
	Content   string    `json:"content"`
	Level     string    `json:"level"`
	Timestamp int64     `json:"timestamp"`
	CreatedAt time.Time `json:"created_at"`
}
