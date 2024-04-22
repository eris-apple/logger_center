package models

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

func TestLog(t *testing.T) *Log {
	t.Helper()

	return &Log{
		Content: "Some log information",
		Level:   "fatal",
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

func TestSession(t *testing.T, userID string) *Session {
	t.Helper()

	return &Session{
		Token:    "my.secret.token",
		UserID:   userID,
		IsActive: true,
	}
}
