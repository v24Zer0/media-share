package image

import "github.com/gin-gonic/gin"

type Handler interface {
	GetImage(ctx *gin.Context)
	CreateImage(ctx *gin.Context)
	DeleteImage(ctx *gin.Context)
}
