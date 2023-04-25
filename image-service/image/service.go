package image

type Service interface {
	GetImage(postID string) ([]byte, error)
	CreateImage(postID string, fileBytes []byte, filename string) error
	DeleteImage(postID string) error
}
