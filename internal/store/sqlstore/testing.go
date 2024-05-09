package sqlstore

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDB(t *testing.T, databaseURL string) (*gorm.DB, func(...string)) {
	t.Helper()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)
	db, connectionErr := gorm.Open(postgres.Open(databaseURL), &gorm.Config{Logger: newLogger})
	if connectionErr != nil {
		t.Fatal(connectionErr)
	}

	sqlDB, sqlErr := db.DB()
	if sqlErr != nil {
		t.Fatal(sqlErr)
	}

	pingErr := sqlDB.Ping()
	if pingErr != nil {
		t.Fatal(pingErr)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))

			err := sqlDB.Close()
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
