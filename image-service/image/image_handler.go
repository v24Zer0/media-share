package image

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/v24Zer0/media-share/image-service/response"
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
		ctx.JSON(400, response.NewErrorResponse("missing postID"))
		return
	}

	b, err := handler.service.GetImage(postID)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, err.Error())
		return
	}

	ctx.Status(200)
	ctx.Writer.Write(b)
}

func (handler ImageHandler) CreateImage(ctx *gin.Context) {
	postID := ctx.PostForm("postID")
	if postID == "" {
		ctx.JSON(400, response.NewErrorResponse("missing postID"))
		return
	}

	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(400, response.NewErrorResponse("missing image"))
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(400, response.NewErrorResponse("failed to open image"))
		return
	}

	// err = ValidateFile(fileHeader.Filename)
	// if err != nil {
	// 	ctx.JSON(400, response.NewErrorResponse("invalid filename format"))
	// 	return
	// }

	fileBytes := []byte{}
	_, err = file.Read(fileBytes)
	defer file.Close()
	if err != nil {
		ctx.JSON(400, response.NewErrorResponse("failed to read image"))
		return
	}

	err = handler.service.CreateImage(postID, fileBytes, fileHeader.Filename)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (handler ImageHandler) DeleteImage(ctx *gin.Context) {
	// Delete image from DB and delete image from file system

	var imageRequest ImageRequest

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&imageRequest)
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	err = handler.service.DeleteImage(imageRequest.PostID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse(err.Error()))
		return
	}
}
