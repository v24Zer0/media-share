package post

import (
	"errors"

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

func (service PostService) GetPosts() (*[]Post, error) {
	return service.repo.GetPosts()
}

func (service PostService) CreatePost(post *Post) error {
	// validate post data:
	// 		check title, createdBy

	// generate post id
	post.ID = service.idProvider.GenerateID()

	err := service.repo.CreatePost(post)
	if err != nil {
		return errors.New("failed to create post")
	}
	return nil
}

func (service PostService) DeletePost(id string) error {
	// validate post data: id, createdBy
	err := service.repo.DeletePost(id)
	if err != nil {
		return errors.New("failed to delete post")
	}
	return nil
}
