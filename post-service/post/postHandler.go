package post

import "github.com/gin-gonic/gin"

type PostHandler struct {
	service Service
}

func NewPostHandler(service Service) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (handler PostHandler) GetPosts(c *gin.Context) {
	handler.service.GetPosts("")
}

func (handler PostHandler) CreatePost(c *gin.Context) {
	post := &Post{}

	handler.service.CreatePost(post)
}

func (handler PostHandler) DeletePost(c *gin.Context) {
	id := ""

	handler.service.DeletePost(id)
}
