package rest_test

import (
	"bytes"
	"encoding/json"
	"github.com/eris-apple/logger_center/internal/config"
	"github.com/eris-apple/logger_center/internal/dto"
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/services"
	"github.com/eris-apple/logger_center/internal/store/teststore"
	"github.com/eris-apple/logger_center/internal/transport/rest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIdentityHandler_SignUp(t *testing.T) {
	testCases := []struct {
		name    string
		body    dto.SignDTO
		err     error
		isValid bool
	}{
		{
			name: "valid",
			body: dto.SignDTO{
				Email:    "mail@example.com",
				Password: "password",
			},
			err:     nil,
			isValid: true,
		},
		{
			name: "invalid email",
			body: dto.SignDTO{
				Email:    "mail.example.com",
				Password: "password",
			},
			err:     config.ErrInvalidEmail,
			isValid: false,
		},
		{
			name: "empty email",
			body: dto.SignDTO{
				Email:    "",
				Password: "password",
			},
			err:     config.ErrBadRequest,
			isValid: false,
		},
		{
			name: "invalid password",
			body: dto.SignDTO{
				Email:    "mail@example.com",
				Password: "123",
			},
			err:     config.ErrInvalidPassword,
			isValid: false,
		},
		{
			name: "empty password",
			body: dto.SignDTO{
				Email:    "mail@example.com",
				Password: "",
			},
			err:     config.ErrBadRequest,
			isValid: false,
		},
		{
			name: "invalid email and password",
			body: dto.SignDTO{
				Email:    "mailexample.com",
				Password: "123",
			},
			err:     config.ErrInvalidEmail,
			isValid: false,
		},
	}

	s := teststore.New()
	ur := s.User()
	sr := s.Session()
	is := services.NewIdentityService(ur, sr)
	ih := rest.IdentityHandler{IdentityService: is}

	gin.SetMode(gin.TestMode)

	apiUrl := "/identity/sign-up"

	router := gin.Default()
	router.POST(apiUrl, ih.SignUp)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.body)
			if err != nil {
				log.Fatalf("impossible to marshall body: %s", err)
			}

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, apiUrl, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(rec, req)

			resBody := &models.HttpResponseError{}
			if err := json.Unmarshal(rec.Body.Bytes(), resBody); err != nil {
				t.Fatalf("failed unmarshal response body")
			}

			if tc.isValid != true {
				assert.Equal(t, resBody.Error, tc.err.Error())
				assert.Equal(t, rec.Code, http.StatusBadRequest)
			}
		})
	}
}

func TestIdentityHandler_SignIn(t *testing.T) {
	testCases := []struct {
		name    string
		body    dto.SignDTO
		err     error
		isValid bool
	}{
		{
			name: "valid",
			body: dto.SignDTO{
				Email:    "mail@example.com",
				Password: "password",
			},
			err:     nil,
			isValid: true,
		},
		{
			name: "invalid email",
			body: dto.SignDTO{
				Email:    "mail.example.com",
				Password: "password",
			},
			err:     config.ErrInvalidEmail,
			isValid: false,
		},
		{
			name: "empty email",
			body: dto.SignDTO{
				Email:    "",
				Password: "password",
			},
			err:     config.ErrBadRequest,
			isValid: false,
		},
		{
			name: "invalid password",
			body: dto.SignDTO{
				Email:    "mail@example.com",
				Password: "123",
			},
			err:     config.ErrInvalidPassword,
			isValid: false,
		},
		{
			name: "empty password",
			body: dto.SignDTO{
				Email:    "mail@example.com",
				Password: "",
			},
			err:     config.ErrBadRequest,
			isValid: false,
		},
		{
			name: "invalid email and password",
			body: dto.SignDTO{
				Email:    "mailexample.com",
				Password: "123",
			},
			err:     config.ErrInvalidEmail,
			isValid: false,
		},
	}

	s := teststore.New()
	ur := s.User()
	sr := s.Session()
	is := services.NewIdentityService(ur, sr)
	ih := rest.IdentityHandler{IdentityService: is}

	gin.SetMode(gin.TestMode)

	apiUrl := "/identity/sign-in"

	router := gin.Default()
	router.POST(apiUrl, ih.SignIn)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.body)
			if err != nil {
				log.Fatalf("impossible to marshall body: %s", err)
			}

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, apiUrl, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(rec, req)

			resBody := &models.HttpResponseError{}
			if err := json.Unmarshal(rec.Body.Bytes(), resBody); err != nil {
				t.Fatalf("failed unmarshal response body")
			}

			if tc.isValid != true {
				assert.Equal(t, resBody.Error, tc.err.Error())
				assert.Equal(t, rec.Code, http.StatusBadRequest)
			}
		})
	}
}
