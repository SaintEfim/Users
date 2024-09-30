package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"Users/config"
	"Users/docs"
	"Users/internal/middleware"
	"Users/internal/models/interfaces"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type Server struct {
	srv     *http.Server
	cfg     *config.Config
	handler interfaces.Handler
	logger  *zap.Logger
}

func NewServer(srv *http.Server, cfg *config.Config, handler interfaces.Handler, logger *zap.Logger) interfaces.Server {
	return &Server{
		srv:     srv,
		cfg:     cfg,
		handler: handler,
		logger:  logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	g := gin.Default()
	g.Use(middleware.LoggingMiddleware(s.logger))

	s.SetGinMode(ctx)
	s.ConfigureSwagger(ctx, g)
	s.handler.ConfigureRoutes(g)

	s.srv.Handler = g

	s.logger.Sugar().Infof("Listening and serving HTTP on %s\n", s.srv.Addr)

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	s.srv.RegisterOnShutdown(cancel)

	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Server) ConfigureSwagger(ctx context.Context, router *gin.Engine) {
	docs.SwaggerInfo.Title = "Users Service API"
	docs.SwaggerInfo.Description = "This is a sample server Users server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *Server) SetGinMode(ctx context.Context) {
	switch s.cfg.EnvironmentVariables.Environment {
	case "development":
		gin.SetMode(gin.DebugMode)
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		log.Printf("Unknown environment: %s, defaulting to 'development'", s.cfg.EnvironmentVariables.Environment)
		gin.SetMode(gin.DebugMode)
	}
}
