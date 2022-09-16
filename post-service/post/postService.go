package post

import "errors"

type PostService struct {
	repo Repo
}

func NewPostService(repo Repo) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (p PostService) GetPosts() {
	p.repo.GetPosts()
}

func (p PostService) CreatePost(post *Post) error {
	return errors.New("failed to create post")
}

func (p PostService) DeletePost(id string) error {
	return errors.New("failed to delete post")
}
