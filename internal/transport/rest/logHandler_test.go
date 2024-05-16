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
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogHandler_Create(t *testing.T) {
	testCases := []struct {
		name    string
		body    dto.CreateLogDTO
		err     error
		isValid bool
	}{
		{
			name: "valid",
			body: dto.CreateLogDTO{
				Content: "some log",
				Level:   "fatal",
			},
			err:     nil,
			isValid: true,
		},
		{
			name: "invalid level",
			body: dto.CreateLogDTO{
				Content: "some log",
				Level:   "some level",
			},
			err:     config.ErrInvalidLogLevel,
			isValid: false,
		},
		{
			name: "empty level",
			body: dto.CreateLogDTO{
				Content: "some log",
			},
			err:     config.ErrInvalidLogLevel,
			isValid: false,
		},
		{
			name: "with valid chainId",
			body: dto.CreateLogDTO{
				ChainID: uuid.NewV4().String(),
				Content: "some log",
				Level:   "fatal",
			},
			err:     nil,
			isValid: true,
		},
		{
			name: "with invalid chainId",
			body: dto.CreateLogDTO{
				ChainID: "123",
				Content: "some log",
				Level:   "fatal",
			},
			err:     config.ErrInvalidLogChainID,
			isValid: false,
		},
	}

	s := teststore.New()
	pr := s.Project()
	lr := s.Log()
	ls := services.NewLogService(lr, pr)
	lh := rest.LogHandler{LogService: ls}

	tp := models.TestProject(t)
	if err := pr.Create(tp); err != nil {
		t.Fatalf("cannot create project")
	}

	gin.SetMode(gin.TestMode)

	apiUrl := "/projects/:project_id/logs"
	fApiUrl := fmt.Sprintf("/projects/%s/logs", tp.ID)

	router := gin.Default()
	router.POST(apiUrl, lh.Create)

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

func TestLogHandler_Update(t *testing.T) {
	testCases := []struct {
		name    string
		body    dto.UpdateLogDTO
		err     error
		isValid bool
	}{
		{
			name: "valid",
			body: dto.UpdateLogDTO{
				Content: "some log",
				Level:   "fatal",
			},
			err:     nil,
			isValid: true,
		},
		{
			name: "invalid level",
			body: dto.UpdateLogDTO{
				Content: "some log",
				Level:   "some level",
			},
			err:     config.ErrInvalidLogLevel,
			isValid: false,
		},
		{
			name: "empty level",
			body: dto.UpdateLogDTO{
				Content: "some log",
			},
			err:     nil,
			isValid: true,
		},
		{
			name: "with valid chainId",
			body: dto.UpdateLogDTO{
				ChainID: uuid.NewV4().String(),
				Content: "some log",
				Level:   "fatal",
			},
			err:     nil,
			isValid: true,
		},
		{
			name: "with invalid chainId",
			body: dto.UpdateLogDTO{
				ChainID: "123",
				Content: "some log",
				Level:   "fatal",
			},
			err:     config.ErrInvalidLogChainID,
			isValid: false,
		},
	}

	s := teststore.New()
	pr := s.Project()
	lr := s.Log()
	ls := services.NewLogService(lr, pr)
	lh := rest.LogHandler{LogService: ls}

	tp := models.TestProject(t)
	if err := pr.Create(tp); err != nil {
		t.Fatalf("cannot create project")
	}

	tl := models.TestLog(t)
	tl.ProjectID = tp.ID

	if err := lr.Create(tl); err != nil {
		t.Fatalf("cannot create log")
	}

	gin.SetMode(gin.TestMode)

	apiUrl := "/projects/:project_id/logs/:log_id"
	fApiUrl := fmt.Sprintf("/projects/%s/logs/%s", tp.ID, tl.ID)

	router := gin.Default()
	router.PUT(apiUrl, lh.Update)

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
