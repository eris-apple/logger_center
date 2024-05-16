package teststore_repository_test

import (
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSessionRepository_Create(t *testing.T) {
	s := teststore.New()

	u1 := models.TestUser(t)
	if err := s.User().Create(u1); err != nil {
		t.Fatal(err)
	}

	ss := models.TestSession(t, u1.ID)

	assert.NoError(t, s.Session().Create(ss))
	assert.NotNil(t, ss.ID)
}

func TestSessionRepository_FindById(t *testing.T) {
	s := teststore.New()

	u1 := models.TestUser(t)
	if err := s.User().Create(u1); err != nil {
		t.Fatal(err)
	}

	ss1 := *models.TestSession(t, u1.ID)
	if err := s.Session().Create(&ss1); err != nil {
		t.Fatal(err)
	}

	ss2, err := s.Session().FindById(ss1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, ss2)
}

func TestSessionRepository_FindByToken(t *testing.T) {
	s := teststore.New()

	u1 := models.TestUser(t)
	if err := s.User().Create(u1); err != nil {
		t.Fatal(err)
	}

	ss1 := models.TestSession(t, u1.ID)

	if err := s.Session().Create(ss1); err != nil {
		t.Fatal(err)
	}

	ss2, err := s.Session().FindByToken(ss1.Token)
	assert.NoError(t, err)
	assert.NotNil(t, ss2)
}
