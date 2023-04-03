package image

type Repo interface {
	GetImage(postID string) (string, error)
	CreateImage(image *Image) (*Image, error)
	DeleteImage(postID string) (*Image, error)
}
