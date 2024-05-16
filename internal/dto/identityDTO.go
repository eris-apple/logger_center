package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type SignDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (sDTO *SignDTO) Validate() error {
	return validation.ValidateStruct(
		sDTO,
		validation.Field(&sDTO.Email, validation.Required, is.Email),
		validation.Field(&sDTO.Password, validation.Required, validation.Length(8, 32)),
	)
}
