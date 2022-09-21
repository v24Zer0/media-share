package mock

import "github.com/v24Zer0/media-share/post-service/post"

type MockService struct {
}

func GetPosts(userID string) (*[]post.Post, error) {
	return &[]post.Post{}, nil
}

func CreatePost(post *post.Post) error {
	return nil
}

func DeletePost(id string) error {
	return nil
}
