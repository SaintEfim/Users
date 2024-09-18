package server

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"Users/config"
	"Users/docs"
	"Users/internal/handler"
	"Users/internal/middleware"
)

type Server struct {
	cfg     *config.Config
	handler *handler.Handler
	logger  *zap.Logger
}

func InitServer(cfg *config.Config, handler *handler.Handler, logger *zap.Logger) *Server {
	return &Server{
		cfg:     cfg,
		handler: handler,
		logger:  logger,
	}
}

func (s *Server) Run() error {
	r := gin.Default()
	r.Use(middleware.LoggingMiddleware(s.logger))

	s.setGinMode()
	s.configureSwagger(r)
	s.handler.ConfigureRoutes(r)

	return r.Run(s.cfg.HTTPServer.Url)
}

func (s *Server) configureSwagger(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Users Service API"
	docs.SwaggerInfo.Description = "This is a sample server Users server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *Server) setGinMode() {
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
