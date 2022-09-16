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

func (h PostHandler) GetPosts(c *gin.Context) {
	h.service.GetPosts()
}

func (h PostHandler) CreatePost(c *gin.Context) {
	post := &Post{}

	h.service.CreatePost(post)
}

func (h PostHandler) DeletePost(c *gin.Context) {
	id := ""

	h.service.DeletePost(id)
}
