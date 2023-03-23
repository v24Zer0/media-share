package image

type Repo interface {
	GetImage(postID string) string
	CreateImage(image *Image) *Image
	DeleteImage(postID string) *Image
}
