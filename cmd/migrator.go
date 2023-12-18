package main

import (
	"example/template/config"
	"example/template/internal/adapters/repo"
	"example/template/pkg/database"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	db, err := database.New(cfg.DB.DSN)

	if err != nil {
		log.Fatalf("Database error: %s", err)
	}

	repo.Migrate(db.DB)
}
