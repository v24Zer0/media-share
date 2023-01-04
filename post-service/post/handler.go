package post

import "github.com/gin-gonic/gin"

type Handler interface {
	GetPosts(c *gin.Context)
	CreatePost(c *gin.Context)
	DeletePost(c *gin.Context)
}
