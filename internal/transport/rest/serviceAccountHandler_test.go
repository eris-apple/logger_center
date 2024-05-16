package rest_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func TestServiceAccountHandler_Create(t *testing.T) {
	testCases := []struct {
		name    string
		body    dto.CreateServiceAccountDTO
		err     error
		isValid bool
	}{
		{
			name: "valid",
			body: dto.CreateServiceAccountDTO{
				Name:        "Test Service Account",
				IsActive:    true,
				Description: "some service account description",
			},
			err:     nil,
			isValid: true,
		},
		{
			name: "invalid name",
			body: dto.CreateServiceAccountDTO{
				IsActive:    true,
				Description: "some service account description",
			},
			err:     config.ErrInvalidServiceAccountName,
			isValid: false,
		},
	}

	s := teststore.New()
	pr := s.Project()
	sar := s.ServiceAccount()
	ps := services.NewProjectService(pr)
	sas := services.NewServiceAccountService(sar, ps)
	ph := rest.ServiceAccountHandler{ServiceAccountService: sas}

	gin.SetMode(gin.TestMode)

	tp := models.TestProject(t)
	if err := pr.Create(tp); err != nil {
		t.Fatalf("cannot create project")
	}

	apiUrl := "/projects/:project_id/service_accounts"
	fApiUrl := fmt.Sprintf("/projects/%s/service_accounts", tp.ID)

	router := gin.Default()
	router.POST(apiUrl, ph.Create)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.body)
			if err != nil {
				log.Fatalf("impossible to marshall body: %s", err)
			}

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, fApiUrl, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(rec, req)

			log.Println(rec.Body.String())

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

func TestServiceAccountHandler_Update(t *testing.T) {
	testCases := []struct {
		name    string
		body    dto.CreateServiceAccountDTO
		err     error
		isValid bool
	}{
		{
			name: "valid",
			body: dto.CreateServiceAccountDTO{
				Name:        "Test Service Account",
				IsActive:    true,
				Description: "some service account description",
			},
			err:     nil,
			isValid: true,
		},
		{
			name: "valid empty name",
			body: dto.CreateServiceAccountDTO{
				IsActive:    true,
				Description: "some service account description",
			},
			err:     nil,
			isValid: true,
		},
	}

	s := teststore.New()
	pr := s.Project()
	sar := s.ServiceAccount()
	ps := services.NewProjectService(pr)
	sas := services.NewServiceAccountService(sar, ps)
	ph := rest.ServiceAccountHandler{ServiceAccountService: sas}

	tp := models.TestProject(t)
	if err := pr.Create(tp); err != nil {
		t.Fatalf("cannot create project")
	}

	gin.SetMode(gin.TestMode)

	apiUrl := "/projects/:project_id/service_accounts"
	fApiUrl := fmt.Sprintf("/projects/%s/service_accounts", tp.ID)

	router := gin.Default()
	router.PUT(apiUrl, ph.Update)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.body)
			if err != nil {
				log.Fatalf("impossible to marshall body: %s", err)
			}

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPut, fApiUrl, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(rec, req)

			log.Println(rec.Body.String())

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
