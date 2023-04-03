package image

import (
	"io/ioutil"
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
	imagePath, err := service.repo.GetImage(postID)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (service ImageService) CreateImage(postID string, file multipart.File, filename string) error {
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
