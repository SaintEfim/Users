package main

import (
	"context"
	"fmt"

	"Users/config"
	"Users/internal/controller"
	"Users/internal/handler"
	"Users/internal/models/interfaces"
	"Users/internal/repository/psql"
	"Users/internal/server"
	"Users/pkg/logger"

	"go.uber.org/fx"
)

func registerServer(lifecycle fx.Lifecycle, srv interfaces.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := srv.Run(); err != nil {
				fmt.Errorf("Failed to start server: %v", err)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

func main() {
	fx.New(
		fx.Provide(func() (*config.Config, error) {
			return config.ReadConfig("config", "yaml", "./config")
		}),
		fx.Provide(
			psql.Connect,
			psql.NewRepository,
			controller.NewController,
			handler.NewHandler,
			logger.InitLogger,
			server.InitServer,
		),
		fx.Invoke(registerServer),
	).Run()
}
