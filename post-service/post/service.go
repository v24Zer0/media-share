package post

type Service interface {
	GetPosts(userID string) (*[]Post, error)
	CreatePost(post *Post) error
	DeletePost(post *Post) error
}
