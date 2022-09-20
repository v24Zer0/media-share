package post_test

import (
	"testing"

	"github.com/v24Zer0/media-share/post-service/mock"
	"github.com/v24Zer0/media-share/post-service/post"
)

func TestServiceGetPosts(t *testing.T) {
	repo := NewMockRepo()
	idProvider := mock.MockIDProvider{}

	service := post.NewPostService(repo, idProvider)

	service.GetPosts()
}

func TestServiceCreatePost(t *testing.T) {
	repo := mock.NewMockRepo()
	idProvider := mock.MockIDProvider{}

	service := post.NewPostService(repo, idProvider)
	post := post.Post{
		ID:        "",
		Title:     "Title 6",
		CreatedAt: "",
		UserID:    "0002",
		CreatedBy: "user2",
	}

	// Should create post and return nil err
	err := service.CreatePost(&post)
	if err != nil {
		t.Fatalf("Failed to create post with id: %s", post.ID)
	}

	// Should fail to create post and return err with message "missing field"
}

func TestServiceDeletePost(t *testing.T) {
	repo := mock.NewMockRepo()
	idProvider := mock.MockIDProvider{}

	service := post.NewPostService(repo, idProvider)
	id := "1"

	err := service.DeletePost(id)
	if err != nil {
		t.Fatalf("Failed to delete post with id: %s", id)
	}
}
