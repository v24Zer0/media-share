package mock

import (
	"errors"

	"github.com/v24Zer0/media-share/post-service/post"
)

type MockRepo struct {
	posts []post.Post
}

func NewMockRepo() *MockRepo {
	return &MockRepo{
		posts: *GetMockPosts(),
	}
}

func (repo MockRepo) GetPosts(userID string) (*[]post.Post, error) {
	var res []post.Post

	for _, p := range repo.posts {
		if p.UserID == userID {
			res = append(res, p)
		}
	}
	return &res, nil
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
