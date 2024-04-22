package sqlstore_test

import (
	"os"
	"testing"
)

var (
	DatabaseURL string
	JWTSecret   string
)

func TestMain(m *testing.M) {
	DatabaseURL = os.Getenv("DATABASE_URL")
	if DatabaseURL == "" {
		DatabaseURL = "postgres://root:password@localhost:5432/logger_center"
	}

	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		JWTSecret = "jwt_secret"
	}

	os.Exit(m.Run())
}
