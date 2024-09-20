package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		logFields := []zap.Field{
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.Path),
		}

		logger.Info("Incoming request",
			logFields...,
		)

		c.Next()

		if len(c.Errors) > 0 {
			logFields = append(logFields, zap.String("error", c.Errors.String()))
			logger.Error("Request completed with errors", logFields...)
		} else {
			logFields = append(logFields,
				zap.Int("status", c.Writer.Status()),
				zap.Duration("duration", time.Since(startTime)),
			)

			logger.Info("Request completed", logFields...)
		}
	}
}
