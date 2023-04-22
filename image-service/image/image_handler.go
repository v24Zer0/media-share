package image

import (
	"encoding/json"

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

	ctx.Status(200)
	ctx.Writer.Write(b)
}

func (handler ImageHandler) CreateImage(ctx *gin.Context) {
	postID := ctx.PostForm("postID")
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.AbortWithStatus(400)
	}

	file, err := fileHeader.Open()
	if err != nil {
		ctx.AbortWithStatus(400)
	}

	err = handler.service.CreateImage(postID, file, fileHeader.Filename)
	if err != nil {
		ctx.AbortWithError(400, err)
	}
}

func (handler ImageHandler) DeleteImage(ctx *gin.Context) {
	// Delete image from DB and delete image from file system

	var imageRequest ImageRequest

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&imageRequest)
	if err != nil {
		ctx.AbortWithStatus(400)
	}

	err = handler.service.DeleteImage(imageRequest.PostID)
	if err != nil {
		ctx.AbortWithError(400, err)
	}
}