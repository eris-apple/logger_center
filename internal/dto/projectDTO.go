package dto

import validation "github.com/go-ozzo/ozzo-validation"

type CreateProjectDTO struct {
	Name        string `json:"name"`
	Prefix      string `json:"prefix"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

func (sDTO *CreateProjectDTO) Validate() error {
	return validation.ValidateStruct(
		sDTO,
		validation.Field(&sDTO.Name, validation.Required),
		validation.Field(&sDTO.Prefix, validation.Required),
		validation.Field(&sDTO.IsActive),
		validation.Field(&sDTO.Description),
	)
}

type UpdateProjectDTO struct {
	Name        string `json:"name"`
	Prefix      string `json:"prefix"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

func (sDTO *UpdateProjectDTO) Validate() error {
	return validation.ValidateStruct(
		sDTO,
		validation.Field(&sDTO.Name),
		validation.Field(&sDTO.Prefix),
		validation.Field(&sDTO.IsActive),
		validation.Field(&sDTO.Description),
	)
}

type FindProjectsDTO struct {
	Name      string `form:"name,omitempty"`
	Prefix    string `form:"prefix,omitempty"`
	IsActive  string `form:"is_active,omitempty"`
	CreatedAt string `form:"created_at,omitempty"`
}
