package post_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/v24Zer0/media-share/post-service/mock"
	"github.com/v24Zer0/media-share/post-service/post"
)

func TestNewPostHandler(t *testing.T) {

}

func TestHandlerGetPosts(t *testing.T) {
	service := mock.NewMockService(mock.NewMockRepo(), mock.MockIDProvider{})
	handler := post.NewPostHandler(service)

	c := gin.Context{}

	handler.GetPosts(&c)
}

func TestHandlerCreatePost(t *testing.T) {
	service := mock.NewMockService(mock.NewMockRepo(), mock.MockIDProvider{})
	handler := post.NewPostHandler(service)

	c := gin.Context{}

	handler.CreatePost(&c)
}

func TestHandlerDeletePost(t *testing.T) {
	service := mock.NewMockService(mock.NewMockRepo(), mock.MockIDProvider{})
	handler := post.NewPostHandler(service)

	c := gin.Context{}

	handler.DeletePost(&c)
}
