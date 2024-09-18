package main

import (
	"log"

	"Users/config"
	"Users/internal/controller"
	"Users/internal/handler"
	"Users/internal/repository/psql"
	"Users/internal/server"
	"Users/pkg"
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

	rep := psql.InitRepository(db, cfg)
	controller := controller.NewController(rep)
	handler := handler.InitHandler(controller)

	logger := logger.InitLogger(cfg)
	srv := server.InitServer(cfg, handler, logger)
	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
