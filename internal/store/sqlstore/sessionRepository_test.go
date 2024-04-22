package sqlstore_test

import (
	"fmt"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSessionRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("sessions", "users")

	s := sqlstore.New(db)

	u1 := models.TestUser(t)
	if err := s.User().Create(u1); err != nil {
		t.Fatal(err)
	}

	ss := models.TestSession(t, u1.ID)

	assert.NoError(t, s.Session().Create(ss))
	assert.NotNil(t, ss.ID)
}

func TestSessionRepository_FindById(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("sessions", "users")

	s := sqlstore.New(db)

	u1 := models.TestUser(t)
	if err := s.User().Create(u1); err != nil {
		t.Fatal(err)
	}

	ss1 := models.TestSession(t, u1.ID)

	fmt.Println(ss1)

	if err := s.Session().Create(ss1); err != nil {
		t.Fatal(err)
	}

	fmt.Println(ss1)

	ss2, err := s.Session().FindById(ss1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, ss2)
}

func TestSessionRepository_FindByToken(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("sessions", "users")

	s := sqlstore.New(db)

	u1 := models.TestUser(t)
	if err := s.User().Create(u1); err != nil {
		t.Fatal(err)
	}

	ss1 := models.TestSession(t, u1.ID)

	fmt.Println(ss1)

	if err := s.Session().Create(ss1); err != nil {
		t.Fatal(err)
	}

	fmt.Println(ss1)

	ss2, err := s.Session().FindByToken(ss1.Token)
	assert.NoError(t, err)
	assert.NotNil(t, ss2)
}
