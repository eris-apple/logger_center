package main

import (
	"github.com/eris-apple/logger_center/internal/app"
	"github.com/eris-apple/logger_center/internal/config"
	"log"
	"time"
)

func main() {
	time.Local = time.UTC
	c := config.NewConfig()

	if err := app.Start(c); err != nil {
		log.Fatal(err)
	}
}
