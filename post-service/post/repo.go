package post

type Repo interface {
	GetPosts() (*[]Post, error)
	CreatePost(post *Post) error
	DeletePost(id string) error
}
