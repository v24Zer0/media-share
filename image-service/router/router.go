package router

import (
	"github.com/gin-gonic/gin"
	"github.com/v24Zer0/media-share/image-service/image"
)

type Router struct{}

func NewRouter(handler image.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/image/:imageID", handler.GetImage)
	r.POST("/image", handler.CreateImage)
	r.DELETE("/image", handler.DeleteImage)

	return r
}
