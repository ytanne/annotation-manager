package main

import (
	"log"

	"github.com/ytanne/annotation-manager/internal/app"
	"github.com/ytanne/annotation-manager/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
