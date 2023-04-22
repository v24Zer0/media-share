package mock

import (
	"errors"

	"github.com/v24Zer0/media-share/image-service/image"
)

type MockRepo struct {
	mockImages []image.Image
}

func NewMockRepo() *MockRepo {
	return &MockRepo{
		mockImages: images,
	}
}

func (repo *MockRepo) GetImage(postID string) (string, error) {
	for _, img := range repo.mockImages {
		if img.PostID == postID {
			return img.Path, nil
		}
	}
	return "", errors.New("no record found")
}

func (repo *MockRepo) CreateImage(image *image.Image) error {
	repo.mockImages = append(repo.mockImages, *image)
	return nil
}

func (repo *MockRepo) DeleteImage(postID string) error {
	newImages := []image.Image{}
	for _, img := range repo.mockImages {
		if img.PostID == postID {
			continue
		}
		newImages = append(newImages, img)
	}
	repo.mockImages = newImages
	return nil
}
