package server

import (
	"context"
	"fmt"
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

func InitServer(srv *http.Server, cfg *config.Config, handler interfaces.Handler, logger *zap.Logger) interfaces.Server {
	return &Server{
		srv:     srv,
		cfg:     cfg,
		handler: handler,
		logger:  logger,
	}
}

func (s *Server) Run() error {
	r := gin.Default()
	r.Use(middleware.LoggingMiddleware(s.logger))

	s.SetGinMode()
	s.ConfigureSwagger(r)
	s.handler.ConfigureRoutes(r)

	address := fmt.Sprintf("%s:%s", s.cfg.HTTPServer.Addr, s.cfg.HTTPServer.Port)

	s.srv = &http.Server{
		Addr:    address,
		Handler: r,
	}

	return r.Run(address)
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.srv.RegisterOnShutdown(cancel)

	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	return nil
}

func (s *Server) ConfigureSwagger(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Users Service API"
	docs.SwaggerInfo.Description = "This is a sample server Users server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *Server) SetGinMode() {
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
