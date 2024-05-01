package models

import (
	"github.com/aetherteam/logger_center/internal/enums"
	"testing"
)

func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Email:    "user@example.org",
		Password: "password",
		Status:   enums.Pending.String(),
	}
}

func TestSession(t *testing.T, userID string) *Session {
	t.Helper()

	return &Session{
		Token:    "my.secret.token",
		UserID:   userID,
		IsActive: true,
	}
}

func TestLog(t *testing.T) *Log {
	t.Helper()

	return &Log{
		Content: "Some log information",
		Level:   enums.Fatal.String(),
	}
}

func TestProject(t *testing.T) *Project {
	t.Helper()

	return &Project{
		Name:     "test project",
		Prefix:   "TP",
		IsActive: true,
	}
}

func TestServiceAccount(t *testing.T) *ServiceAccount {
	t.Helper()

	return &ServiceAccount{
		Name:     "test account",
		Secret:   "some secret",
		IsActive: true,
	}
}
