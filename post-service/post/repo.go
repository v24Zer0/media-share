package post

type Repo interface {
	GetPosts()
	CreatePost(post *Post)
	DeletePost(id string)
}
