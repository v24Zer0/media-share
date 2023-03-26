package image

import (
	"mime/multipart"
	"os"

	"github.com/v24Zer0/media-share/image-service/util"
)

type ImageService struct {
	repo       Repo
	idProvider util.IDProvider
}

func NewImageService(repo Repo, idProvider util.IDProvider) *ImageService {
	return &ImageService{
		repo:       repo,
		idProvider: idProvider,
	}
}

func (service ImageService) GetImage(postID string) ([]byte, error) {
	// Return image as bytes written to conttext Writer
	// b, err := ioutil.ReadFileutil
	// if err != nil {
	// 	panic(err)
	// }

	// ctx.Writer.Write(b)
	return []byte{}, nil
}

func (service ImageService) CreateImage(postID string, file multipart.File, filename string) error {
	// Verify file adheres to format (no malicious filename)
	// Verify file is correct format
	err := ValidateFile(filename)
	if err != nil {
		return err
	}

	id := service.idProvider.GenerateID()
	path := service.idProvider.GenerateFileID()

	image := &Image{
		ID:     id,
		Path:   path,
		PostID: postID,
	}

	fileBytes := []byte{}
	file.Read(fileBytes)
	defer file.Close()

	service.repo.CreateImage(image)

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	f.Write(fileBytes)
	f.Close()

	return nil
}

func (service ImageService) DeleteImage(postID string) error {
	return nil
}

func ValidateFile(filename string) error {
	return nil
}

func GetFileExtension(filename string) (string, error) {
	return "", nil
}
