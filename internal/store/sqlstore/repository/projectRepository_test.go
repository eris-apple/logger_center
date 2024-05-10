package repository_test

import (
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store/sqlstore"
	"github.com/eris-apple/logger_center/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("projects")

	s := sqlstore.New(db)
	p := models.TestProject(t)

	assert.NoError(t, s.Project().Create(p))
	assert.NotNil(t, p.ID)
}

func TestProjectRepository_FindAll(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("projects")

	s := sqlstore.New(db)
	p1 := models.TestProject(t)

	if err := s.Project().Create(p1); err != nil {
		t.Fatal(err)
	}

	p2, err := s.Project().FindAll(&utils.Filter{}, make(map[string]interface{}))
	assert.NoError(t, err)
	assert.NotNil(t, p2)
}

func TestProjectRepository_FindById(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("projects")

	s := sqlstore.New(db)
	p1 := models.TestProject(t)

	if err := s.Project().Create(p1); err != nil {
		t.Fatal(err)
	}

	p2, err := s.Project().FindById(p1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, p2)
}

func TestProjectRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("projects")

	s := sqlstore.New(db)
	p := models.TestProject(t)

	p.IsActive = false

	err := s.Project().Update(p)
	assert.NoError(t, err)
	assert.NotNil(t, p)
}
