package post_test

import (
	"errors"
	"testing"

	"github.com/v24Zer0/media-share/post-service/mock"
	"github.com/v24Zer0/media-share/post-service/post"
)

type MockRepo struct {
	posts []post.Post
}

func NewMockRepo() *MockRepo {
	return &MockRepo{
		posts: *mock.GetMockPosts(),
	}
}

func (repo MockRepo) GetPosts(userID string) (*[]post.Post, error) {
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

func TestRepoGetPosts(t *testing.T) {

}

func TestRepoCreatePost(t *testing.T) {
	repo := NewMockRepo()
	post := post.Post{
		ID:        "6",
		Title:     "Title 1",
		CreatedAt: "",
		UserID:    "0001",
		CreatedBy: "user1",
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

func TestRepoDeletePost(t *testing.T) {

}
