package image

type Repo interface {
	GetImage(postID string) (string, error)
	CreateImage(image *Image) error
	DeleteImage(postID string) (string, error)
}
