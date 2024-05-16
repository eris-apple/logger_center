package teststore_repository_test

import (
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store/teststore"
	"github.com/eris-apple/logger_center/internal/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogRepository_Create(t *testing.T) {
	s := teststore.New()

	p := models.TestProject(t)
	assert.NoError(t, s.Project().Create(p))
	assert.NotNil(t, p.ID)

	l := models.TestLog(t)
	l.ProjectID = p.ID

	assert.NoError(t, s.Log().Create(l))
	assert.NotNil(t, l.ID)
}

func TestLogRepository_FindAll(t *testing.T) {
	s := teststore.New()

	p := models.TestProject(t)
	assert.NoError(t, s.Project().Create(p))
	assert.NotNil(t, p.ID)

	l := models.TestLog(t)
	l.ProjectID = p.ID

	assert.NoError(t, s.Log().Create(l))
	assert.NotNil(t, l.ID)

	logs, err := s.Log().FindAll(l.ProjectID, &utils.Filter{})
	assert.NoError(t, err)
	assert.NotNil(t, logs)
}

func TestLogRepository_FindById(t *testing.T) {
	s := teststore.New()

	p := models.TestProject(t)
	assert.NoError(t, s.Project().Create(p))
	assert.NotNil(t, p.ID)

	l := models.TestLog(t)
	l.ProjectID = p.ID

	assert.NoError(t, s.Log().Create(l))
	assert.NotNil(t, l.ID)

	log, err := s.Log().FindById(l.ID)
	assert.NoError(t, err)
	assert.NotNil(t, log)
}

func TestLogRepository_FindByChainId(t *testing.T) {
	s := teststore.New()

	p := models.TestProject(t)
	assert.NoError(t, s.Project().Create(p))
	assert.NotNil(t, p.ID)

	l := models.TestLog(t)
	l.ProjectID = p.ID
	l.ChainID = uuid.NewV4().String()

	assert.NoError(t, s.Log().Create(l))
	assert.NotNil(t, l.ID)

	log, err := s.Log().FindByChainId(l.ChainID, &utils.Filter{})
	assert.NoError(t, err)
	assert.NotNil(t, log)
}

func TestLogRepository_Update(t *testing.T) {
	s := teststore.New()

	p := models.TestProject(t)
	assert.NoError(t, s.Project().Create(p))
	assert.NotNil(t, p.ID)

	l := models.TestLog(t)
	l.ProjectID = p.ID
	l.ChainID = uuid.NewV4().String()

	assert.NoError(t, s.Log().Create(l))
	assert.NotNil(t, l.ID)

	l.Content = "updated some log"

	err := s.Log().Update(l)
	assert.NoError(t, err)
	assert.NotNil(t, l)
}
