package models_test

import (
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *models.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *models.User {
				return models.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *models.User {
				return &models.User{
					Email:    "",
					Password: "password",
				}
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *models.User {
				return &models.User{
					Email:    "invalid",
					Password: "password",
				}
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *models.User {
				return &models.User{
					Email:    "test@email.com",
					Password: "",
				}
			},
			isValid: false,
		},
		{
			name: "invalid password",
			u: func() *models.User {
				return &models.User{
					Email:    "test@email.com",
					Password: "123",
				}
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			}
		})
	}
}
