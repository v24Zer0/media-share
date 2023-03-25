package image

import "mime/multipart"

type Service interface {
	GetImage(postID string) ([]byte, error)
	CreateImage(postID string, fileHeader *multipart.FileHeader) error
	DeleteImage(postID string) error
}
