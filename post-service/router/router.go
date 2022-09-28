package router

import (
	"github.com/gin-gonic/gin"
	"github.com/v24Zer0/media-share/post-service/post"
)

func NewRouter(handler post.Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/post/:userID", handler.GetPosts)
	r.POST("/post", handler.CreatePost)
	r.DELETE("/post", handler.DeletePost)

	return r
}
