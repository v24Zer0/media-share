package image

import (
	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	service *Service
}

func NewImageHandler(service *Service) *ImageHandler {
	return &ImageHandler{
		service: service,
	}
}

func (handler ImageHandler) GetImage(ctx *gin.Context) {
	// Return image as bytes written to conttext Writer
	// b, err := ioutil.ReadFile("")
	// if err != nil {
	// 	panic(err)
	// }

	// ctx.Writer.Write(b)
}

func (handler ImageHandler) CreateImage(ctx *gin.Context) {
	// Create image in DB and store image in created path
}

func (handler ImageHandler) DeleteImage(ctx *gin.Context) {
	// Delete image from DB and delete image from file system
}
