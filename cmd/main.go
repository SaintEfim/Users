package main

import (
	"context"
	"fmt"

	"Users/config"
	"Users/internal/controller"
	"Users/internal/handler"
	"Users/internal/http_server"
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
				return fmt.Errorf("failed to start server: %w", err)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := srv.Stop(); err != nil {
				return fmt.Errorf("failed to stop server: %w", err)
			}
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
			http_server.InitHTTPServer,
			server.InitServer,
		),
		fx.Invoke(registerServer),
	).Run()
}
