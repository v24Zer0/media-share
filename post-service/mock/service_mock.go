package mock

import (
	"errors"

	"github.com/v24Zer0/media-share/post-service/post"
	"github.com/v24Zer0/media-share/post-service/util"
)

type MockService struct {
	repo       post.Repo
	idProvider util.IDProvider
}

func NewMockService(repo post.Repo, idProvider util.IDProvider) *MockService {
	return &MockService{
		repo:       repo,
		idProvider: idProvider,
	}
}

func (service MockService) GetPosts(userID string) (*[]post.Post, error) {
	if userID == "" {
		return &[]post.Post{}, errors.New("missing userID")
	}
	return service.repo.GetPosts(userID)
}

func (service MockService) CreatePost(post *post.Post) error {
	if post.Title == "" || post.UserID == "" || post.CreatedBy == "" {
		return errors.New("missing field")
	}

	return service.repo.CreatePost(post)
}

func (service MockService) DeletePost(id string) error {
	if id == "" {
		return errors.New("missing field")
	}

	return service.repo.DeletePost(id)
}
