package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run(ctx context.Context) error
	Stop(ctx context.Context) error
	ConfigureSwagger(ctx context.Context, router *gin.Engine)
	SetGinMode(ctx context.Context)
}
