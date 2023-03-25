package image

import (
	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	service Service
}

func NewImageHandler(service Service) *ImageHandler {
	return &ImageHandler{
		service: service,
	}
}

func (handler ImageHandler) GetImage(ctx *gin.Context) {
	postID, ok := ctx.Params.Get("postID")
	if !ok {
		ctx.AbortWithStatus(400)
	}

	b, err := handler.service.GetImage(postID)
	if err != nil {
		ctx.AbortWithStatus(400)
	}

	ctx.Writer.Write(b)
}

func (handler ImageHandler) CreateImage(ctx *gin.Context) {
	// Create image in DB and store image in created path
	postID := ctx.PostForm("postID")
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.AbortWithStatus(400)
	}

	file, err := fileHeader.Open()
	if err != nil {
		ctx.AbortWithStatus(400)
	}

	handler.service.CreateImage(postID, file)
}

func (handler ImageHandler) DeleteImage(ctx *gin.Context) {
	// Delete image from DB and delete image from file system
}
