package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		logger.Info("Incoming request",
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.Path),
		)

		c.Next()

		duration := time.Since(startTime)
		logger.Info("Request completed",
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", duration),
		)
	}
}
