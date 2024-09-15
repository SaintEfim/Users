package main

import (
	"log"

	"Users/config"
	"Users/internal/handler"
	"Users/internal/repository/psql"
)

func main() {
	cfg, err := config.ReadConfig("config", "yaml", "./config")
	if err != nil {
		log.Fatalf("Failed to read configuration: %v", err)
	}

	db, err := psql.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	rep, err := psql.InitRepository(db, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	handler := handler.InitServer(cfg, rep)
	if err := handler.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
