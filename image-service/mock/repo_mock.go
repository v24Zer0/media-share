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
	if image.PostID == "post_00" {
		return errors.New("invalid postID")
	}

	repo.mockImages = append(repo.mockImages, *image)
	return nil
}

func (repo *MockRepo) DeleteImage(postID string) (string, error) {
	for _, img := range repo.mockImages {
		if img.PostID == postID {
			return img.Path, nil
		}
	}
	return "", errors.New("no record found")
}
