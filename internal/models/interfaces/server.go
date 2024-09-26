package interfaces

import "github.com/gin-gonic/gin"

type Server interface {
	Run() error
	Stop() error
	ConfigureSwagger(router *gin.Engine)
	SetGinMode()
}
