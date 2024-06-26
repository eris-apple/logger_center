package repository_test

import (
	"github.com/eris-apple/logger_center/internal/config"
	"os"
	"testing"
)

var (
	DatabaseURL string
	JWTSecret   string
)

func TestMain(m *testing.M) {
	DatabaseURL = os.Getenv(config.EnvKeyDatabaseUrl)
	if DatabaseURL == "" {
		DatabaseURL = "postgres://root:password@localhost:5432/logger_center"
	}

	JWTSecret = os.Getenv(config.EnvKeyJWTSecret)
	if JWTSecret == "" {
		JWTSecret = "jwt_secret"
	}

	os.Exit(m.Run())
}
