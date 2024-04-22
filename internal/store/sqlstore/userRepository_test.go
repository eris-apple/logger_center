package sqlstore_test

import (
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/aetherteam/logger_center/internal/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := models.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.ID)
}

func TestUserRepository_FindAll(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := models.TestUser(t)

	if err := s.User().Create(u1); err != nil {
		t.Fatal(err)
	}

	users, err := s.User().FindAll()
	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestUserRepository_FindById(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := models.TestUser(t)

	if err := s.User().Create(u1); err != nil {
		t.Fatal(err)
	}

	u2, err := s.User().FindById(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
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
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := models.TestUser(t)

	u.Email = "email2@example.com"

	result, err := s.User().Update(u)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
