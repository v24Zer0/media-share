package image

import "github.com/v24Zer0/media-share/image-service/util"

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

func (service ImageService) CreateImage(postID string) error {
	id := service.idProvider.GenerateID()
	// path := create new image path
	image := &Image{
		ID:     id,
		Path:   "",
		PostID: postID,
	}

	service.repo.CreateImage(image)
	return nil
}

func (service ImageService) DeleteImage(postID string) error {
	return nil
}
