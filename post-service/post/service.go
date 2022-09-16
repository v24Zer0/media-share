package post

type Service interface {
	GetPosts()
	CreatePost(post *Post) error
	DeletePost(id string) error
}
