package image

import (
	"errors"
	"fmt"
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
	if id == "" {
		return errors.New("error creating image id")
	}

	extension := GetFileExtension(filename)
	path := CreateFilePath(id, extension)

	image := &Image{
		ID:     id,
		Path:   path,
		PostID: postID,
	}

	fileBytes := []byte{}
	file.Read(fileBytes)
	defer file.Close()

	err = service.repo.CreateImage(image)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	f.Write(fileBytes)
	f.Close()

	return nil
}

func (service ImageService) DeleteImage(postID string) error {
	err := service.repo.DeleteImage(postID)
	if err != nil {
		return err
	}

	err = os.Remove("")
	if err != nil {
		return errors.New("error deleting file")
	}

	return nil
}

func CreateFilePath(id string, extension string) string {
	return fmt.Sprintf("%s.%s", id, extension)
}

func ValidateFile(filename string) error {
	return nil
}

func GetFileExtension(filename string) string {
	return ""
}
