package image

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

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
		return nil, errors.New("failed to read file")
	}

	return b, nil
}

func (service ImageService) CreateImage(postID string, fileBytes []byte, filename string) error {
	id := service.idProvider.GenerateID()
	if id == "" {
		return errors.New("error creating image id")
	}

	extension := GetFileExtension(filename)
	if extension == "" {
		return errors.New("failed to get extension")
	}

	path := CreateFilePath(id, extension)

	image := &Image{
		ID:     id,
		Path:   path,
		PostID: postID,
	}

	err := service.repo.CreateImage(image)
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

	err = os.Remove("path_to_image")
	if err != nil {
		return errors.New("error deleting file")
	}

	return nil
}

func CreateFilePath(id string, extension string) string {
	return fmt.Sprintf("upload/images/%s.%s", id, extension)
}

func ValidateFile(filename string) error {
	regex, err := regexp.Compile(`^([a-zA-Z0-9\s_.\\/\-\(\):])+(.jpe?g|.png|.pdf)$`)
	if err != nil {
		return errors.New("failed to compile regex")
	}

	ok := regex.MatchString(filename)
	if !ok {
		return errors.New("invalid filename")
	}
	return nil
}

func GetFileExtension(filename string) string {
	splitString := strings.Split(filename, ".")
	if len(splitString) == 1 {
		return ""
	}

	extension := splitString[len(splitString)-1]
	return extension
}
