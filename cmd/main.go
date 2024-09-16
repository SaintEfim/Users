package main

import (
	"Users/internal/controller"
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

	rep := psql.InitRepository(db, cfg)
	userService := controller.InitController(rep)
	userHandler := handler.InitHandler(userService)

	srv := server.InitServer(cfg, userHandler)
	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
