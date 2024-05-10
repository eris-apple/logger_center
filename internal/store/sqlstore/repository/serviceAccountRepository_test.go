package repository_test

import (
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store/sqlstore"
	"github.com/eris-apple/logger_center/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceAccountRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("service_accounts", "projects")

	s := sqlstore.New(db)

	p := models.TestProject(t)
	assert.NoError(t, s.Project().Create(p))
	assert.NotNil(t, p.ID)

	sa := models.TestServiceAccount(t)
	sa.ProjectID = p.ID

	assert.NoError(t, s.ServiceAccount().Create(sa))
	assert.NotNil(t, sa.ID)
}

func TestServiceAccountRepository_FindAll(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("service_accounts", "projects")

	s := sqlstore.New(db)

	p := models.TestProject(t)
	assert.NoError(t, s.Project().Create(p))
	assert.NotNil(t, p.ID)

	sa := models.TestServiceAccount(t)
	sa.ProjectID = p.ID

	assert.NoError(t, s.ServiceAccount().Create(sa))
	assert.NotNil(t, sa.ID)

	logs, err := s.ServiceAccount().FindAll(p.ID, &utils.Filter{})
	assert.NoError(t, err)
	assert.NotNil(t, logs)
}

func TestServiceAccountRepository_FindById(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("service_accounts", "projects")

	s := sqlstore.New(db)

	p := models.TestProject(t)
	assert.NoError(t, s.Project().Create(p))
	assert.NotNil(t, p.ID)

	sa := models.TestServiceAccount(t)
	sa.ProjectID = p.ID

	assert.NoError(t, s.ServiceAccount().Create(sa))
	assert.NotNil(t, sa.ID)

	result, err := s.ServiceAccount().FindById(sa.ID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestServiceAccountRepository_FindBySecret(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("service_accounts", "projects")

	s := sqlstore.New(db)

	p := models.TestProject(t)
	assert.NoError(t, s.Project().Create(p))
	assert.NotNil(t, p.ID)

	sa := models.TestServiceAccount(t)
	sa.ProjectID = p.ID

	assert.NoError(t, s.ServiceAccount().Create(sa))
	assert.NotNil(t, sa.ID)

	result, err := s.ServiceAccount().FindBySecret(sa.Secret)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestServiceAccountRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("service_accounts", "projects")

	s := sqlstore.New(db)

	p := models.TestProject(t)
	assert.NoError(t, s.Project().Create(p))
	assert.NotNil(t, p.ID)

	sa := models.TestServiceAccount(t)
	sa.ProjectID = p.ID

	assert.NoError(t, s.ServiceAccount().Create(sa))
	assert.NotNil(t, sa.ID)

	sa.Secret = "updated secret"

	assert.NoError(t, s.ServiceAccount().Update(sa))
	assert.NotNil(t, sa)
}
