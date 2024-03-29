package post

import (
	"errors"
	"time"

	"github.com/v24Zer0/media-share/post-service/util"
)

type PostService struct {
	repo       Repo
	idProvider util.IDProvider
}

func NewPostService(repo Repo, idProvider util.IDProvider) *PostService {
	return &PostService{
		repo:       repo,
		idProvider: idProvider,
	}
}

func (service PostService) GetPosts(userID string) (*[]Post, error) {
	if userID == "" {
		return &[]Post{}, errors.New("missing userID")
	}
	return service.repo.GetPosts(userID)
}

func (service PostService) CreatePost(post *Post) error {
	// validate post data: title, userID, createdBy
	if post.Title == "" || post.UserID == "" || post.CreatedBy == "" {
		return errors.New("missing field")
	}

	if post.CreatedAt == "" {
		post.CreatedAt = time.Now().String()
	}

	// generate post id
	post.ID = service.idProvider.GenerateID()

	err := service.repo.CreatePost(post)
	if err != nil {
		return errors.New("failed to create post")
	}
	return nil
}

func (service PostService) DeletePost(post *Post) error {
	// validate post data: id
	if post.ID == "" || post.UserID == "" {
		return errors.New("missing field")
	}

	err := service.repo.DeletePost(post)
	if err != nil {
		return errors.New("failed to delete post")
	}
	return nil
}
