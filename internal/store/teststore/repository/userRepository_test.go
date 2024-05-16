package teststore_repository_test

import (
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store"
	"github.com/eris-apple/logger_center/internal/store/teststore"
	"github.com/eris-apple/logger_center/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := models.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.ID)
}

func TestUserRepository_FindAll(t *testing.T) {
	s := teststore.New()
	u1 := models.TestUser(t)

	if err := s.User().Create(u1); err != nil {
		t.Fatal(err)
	}

	users, err := s.User().FindAll(&utils.Filter{}, make(map[string]interface{}))
	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestUserRepository_FindById(t *testing.T) {
	s := teststore.New()
	u1 := *models.TestUser(t)

	if err := s.User().Create(&u1); err != nil {
		t.Fatal(err)
	}

	u2, err := s.User().FindById(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	u1 := models.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	if err := s.User().Create(u1); err != nil {
		t.Fatal(err)
	}

	u2, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_Update(t *testing.T) {
	s := teststore.New()
	u := *models.TestUser(t)

	u.Email = "email2@example.com"

	if err := s.User().Create(&u); err != nil {
		t.Fatal(err)
	}

	err := s.User().Update(&u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
