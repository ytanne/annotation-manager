package app

import (
	"log"

	"github.com/ytanne/annotation-manager/internal/config"
	v1 "github.com/ytanne/annotation-manager/internal/controller/http/v1"
	"github.com/ytanne/annotation-manager/internal/repository/database"
	"github.com/ytanne/annotation-manager/internal/usecase"
	"github.com/ytanne/annotation-manager/pkg/sqlite"
)

func Run(cfg config.Config) {
	db, err := sqlite.NewClient(cfg.DbPath)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repo := database.NewRepository(db)

	u := usecase.NewUsecase(repo)

	c, err := v1.NewController(cfg.AuthConfig, u)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}
