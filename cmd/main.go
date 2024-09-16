package main

import (
	"Users/internal/server"
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

	handler := handler.InitHandler(rep)

	srv := server.InitServer(cfg, handler)
	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
