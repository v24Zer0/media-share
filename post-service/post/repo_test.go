package post_test

import (
	"errors"
	"testing"

	"github.com/v24Zer0/media-share/post-service/mock"
	"github.com/v24Zer0/media-share/post-service/post"
	"github.com/v24Zer0/media-share/post-service/user"
)

type MockRepo struct {
	posts []post.Post
}

func NewMockRepo() *MockRepo {
	return &MockRepo{
		posts: *mock.GetMockPosts(),
	}
}

func (repo MockRepo) GetPosts() (*[]post.Post, error) {
	return &repo.posts, nil
}

func (repo MockRepo) CreatePost(post *post.Post) error {
	for _, p := range repo.posts {
		if p.ID == post.ID {
			return errors.New("failed to create post")
		}
	}

	return nil
}

func (repo MockRepo) DeletePost(id string) error {
	for _, p := range repo.posts {
		if p.ID == id {
			return nil
		}
	}

	return errors.New("failed to delete post")
}

func TestNewPostRepo(t *testing.T) {
}

func TestGetPosts(t *testing.T) {

}

func TestCreatePost(t *testing.T) {
	repo := NewMockRepo()
	post := post.Post{
		ID:        "6",
		Title:     "",
		CreatedAt: "",
		CreatedBy: user.User{
			ID:       "0001",
			Username: "user1",
		},
		ImageID: "",
	}

	// Valid CreatePost()
	err := repo.CreatePost(&post)
	if err != nil {
		t.Fatal("Failed to create valid post")
	}

	// post.ID = "1"

	// Invalid CreatePost()
	// err = repo.CreatePost(&post)
}

func TestDeletePost(t *testing.T) {

}
