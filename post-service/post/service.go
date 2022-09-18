package post

type Service interface {
	GetPosts() (*[]Post, error)
	CreatePost(post *Post) error
	DeletePost(id string) error
}
