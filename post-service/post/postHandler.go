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
	userID, ok := c.Params.Get("userID")
	if !ok {
		c.AbortWithStatus(400)
		return
	}

	posts, err := handler.service.GetPosts(userID)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	c.JSON(200, posts)
}

func (handler PostHandler) CreatePost(c *gin.Context) {
	post := Post{}
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	err = handler.service.CreatePost(&post)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	c.Status(201)
}

func (handler PostHandler) DeletePost(c *gin.Context) {
	post := Post{}
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	err = handler.service.DeletePost(&post)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	c.Status(200)
}
