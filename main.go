package main

import (
	"github.com/aetherteam/logger_center/internal/app"
	"github.com/aetherteam/logger_center/internal/config"
	"log"
)

func main() {
	c := config.NewConfig()

	if err := app.Start(c); err != nil {
		log.Fatal(err)
	}

}
