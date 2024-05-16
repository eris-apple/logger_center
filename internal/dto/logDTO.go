package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UpdateLogDTO struct {
	ChainID   string `json:"chain_id"`
	ProjectID string `json:"project_id"`
	Title     string `json:"title"`
	Error     string `json:"error"`
	Params    string `json:"params"`
	Content   string `json:"content"`
	Timestamp int64  `json:"Timestamp"`
	Level     string `json:"level"`
}

type CreateLogDTO struct {
	ChainID   string `json:"chain_id"`
	Title     string `json:"title"`
	Error     string `json:"error"`
	Params    string `json:"params"`
	Content   string `json:"content"`
	Timestamp int64  `json:"Timestamp"`
	Level     string `json:"level"`
}

func (clDTO *CreateLogDTO) Validate() error {
	return validation.ValidateStruct(
		clDTO,
		validation.Field(&clDTO.ChainID, is.UUID),
		validation.Field(&clDTO.Level, validation.Required, validation.In("info", "alert", "debug", "warning", "error", "fatal")),
	)
}

func (clDTO *UpdateLogDTO) Validate() error {
	return validation.ValidateStruct(
		clDTO,
		validation.Field(&clDTO.ChainID, is.UUID),
		validation.Field(&clDTO.Level, validation.In("info", "alert", "debug", "warning", "error", "fatal")),
	)
}
