package dto

import validation "github.com/go-ozzo/ozzo-validation"

type CreateServiceAccountDTO struct {
	IsActive    bool   `json:"is_active"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateServiceAccountDTO struct {
	IsActive    bool   `json:"is_active"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (casDTO *CreateServiceAccountDTO) Validate() error {
	return validation.ValidateStruct(
		casDTO,
		validation.Field(&casDTO.Name, validation.Required),
	)
}

func (uasDTO *UpdateServiceAccountDTO) Validate() error {
	return validation.ValidateStruct(
		uasDTO,
	)
}
