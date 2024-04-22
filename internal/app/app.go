package app

import (
	"database/sql"
	"github.com/aetherteam/logger_center/internal/config"
	"github.com/aetherteam/logger_center/internal/store/sqlstore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Start(config *config.Config) error {
	db, sqlDB, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(sqlDB)

	store := sqlstore.New(db)
	srv := newServer(config, store)

	log.Println("Server has been started", config.BindAddress)
	return http.ListenAndServe(config.BindAddress, srv)
}

func newDB(databaseURL string) (*gorm.DB, *sql.DB, error) {
	db, connectionErr := gorm.Open(postgres.Open(databaseURL), &gorm.Config{TranslateError: true})
	if connectionErr != nil {
		return nil, nil, connectionErr
	}

	sqlDB, sqlErr := db.DB()
	if sqlErr != nil {
		return nil, nil, sqlErr
	}

	pingErr := sqlDB.Ping()
	if pingErr != nil {
		return nil, nil, pingErr
	}

	return db, sqlDB, nil
}
