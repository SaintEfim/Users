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

func registerServer(ctx context.Context, lifecycle fx.Lifecycle, srv interfaces.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := srv.Run(ctx); err != nil {
				return fmt.Errorf("failed to start server: %w", err)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := srv.Stop(ctx); err != nil {
				return fmt.Errorf("failed to stop server: %w", err)
			}
			return nil
		},
	})
}

func registerDependencies() fx.Option {
	return fx.Provide(
		psql.Connect,
		psql.NewRepository,
		controller.NewController,
		handler.NewHandler,
		logger.InitLogger,
		server.NewHTTPServer,
		server.NewServer,
	)
}

func main() {
	fx.New(
		fx.Provide(func() context.Context {
			return context.Background()
		}),
		fx.Provide(func() (*config.Config, error) {
			return config.ReadConfig("config", "yaml", "./config")
		}),
		registerDependencies(),
		fx.Invoke(registerServer),
	).Run()
}
