package middleware

import (
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		logFields := []zap.Field{
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.String("hader", fmt.Sprintf("%v", c.Request.Header)),
			zap.String("request_body", readRequestBody(c.Request.Body)),
			zap.String("query_parameters", c.Request.URL.Query().Encode()),
			zap.String("size_request", fmt.Sprintf("%d", c.Writer.Size())),
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

func readRequestBody(body io.ReadCloser) string {
	if body == nil {
		return ""
	}
	defer body.Close()
	buf, err := io.ReadAll(body)
	if err != nil {
		return ""
	}
	return string(buf)
}
