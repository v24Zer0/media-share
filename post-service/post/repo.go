package post

type Repo interface {
	GetPosts(userID string) (*[]Post, error)
	CreatePost(post *Post) error
	DeletePost(id string) error
}
