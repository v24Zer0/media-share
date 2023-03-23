package image

type Service interface {
	GetImage(postID string) ([]byte, error)
	CreateImage(postID string) error
	DeleteImage(postID string) error
}
