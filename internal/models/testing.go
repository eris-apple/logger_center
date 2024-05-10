package models

import (
	"github.com/eris-apple/logger_center/internal/enums"
	"testing"
	"time"
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
		Content:   "Some log information",
		Level:     enums.Fatal.String(),
		Timestamp: time.Now().Unix(),
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
