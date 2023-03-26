package image

import "mime/multipart"

type Service interface {
	GetImage(postID string) ([]byte, error)
	CreateImage(postID string, file multipart.File, filename string) error
	DeleteImage(postID string) error
}
