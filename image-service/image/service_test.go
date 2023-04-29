package image_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/v24Zer0/media-share/image-service/image"
	"github.com/v24Zer0/media-share/image-service/mock"
	"github.com/v24Zer0/media-share/image-service/util"
)

func TestService_GetImage(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &util.DefaultProvider{})

	img, err := service.GetImage("post_01")
	if err != nil {
		t.Log(err.Error())
		t.Fatal("Failed to retrieve valid image")
	}

	actualImg, err := ioutil.ReadFile("upload/images/700508e7-b810-496e-b7e3-d8d1c6fcf83e.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}

	ok := bytes.Equal(img, actualImg)
	if !ok {
		t.Fatal("Image bytes not equal to actual image")
	}
}

func TestService_GetImageWithUnknownPostImage(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &util.DefaultProvider{})

	_, err := service.GetImage("post_04")
	if err == nil {
		t.Fatal("No error when retrieving unknown post image")
	}
}

func TestService_GetImageWithInvalidPathToImage(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &util.DefaultProvider{})

	_, err := service.GetImage("post_00")
	if err == nil {
		t.Fatal("No error when retrieving image with invalid path")
	}
}

func TestService_CreateImage(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &mock.MockIDProvider{})

	err := service.CreateImage("post_04", mockImage, "mock_image_00.jpg")
	if err != nil {
		t.Fatal("Error when creating valid image")
	}

	// Validate created image exists in directory
	_, err = os.Stat("upload/images/mock-image-id-00.jpg")
	if err != nil {
		t.Fatal("Image not found in directory")
	}

	RestoreStateAfterCreate()
}

func TestService_CreateImageWithInvalidPostID(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &mock.MockIDProvider{})

	err := service.CreateImage("post_00", mockImage, "mock_image_00.jpg")
	if err == nil {
		t.Fatal("Failed to return error for image with invalid postID")
	}

	RestoreStateAfterCreate()
}

func TestService_CreateImageWithInvalidFilename(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &mock.MockIDProvider{})

	err := service.CreateImage("post_01", mockImage, "mock_image_00.pdf")
	if err == nil {
		t.Fatal("Failed to return error for image with invalid filaname")
	}

	RestoreStateAfterCreate()
}

func TestService_CreateImageWithInvalidFileExtension(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &mock.MockIDProvider{})

	err := service.CreateImage("post_04", mockImage, "mock_image_00")
	if err == nil {
		t.Fatal("Failed to return error for image with invalid file extension")
	}

	RestoreStateAfterCreate()
}

func TestService_DeleteImage(t *testing.T) {
	// service := image.NewImageService(mock.NewMockRepo(), &util.DefaultProvider{})

	// err := service.DeleteImage("post_01")
	// if err != nil {
	// 	t.Fatal("Failed to delete valid image")
	// }
}

func TestService_DeleteImageWithInvalidPostID(t *testing.T) {
	// service := image.NewImageService(mock.NewMockRepo(), &util.DefaultProvider{})

	// err := service.DeleteImage("post_04")
	// if err == nil {
	// 	t.Fatal("Failed to return error for invalid postID")
	// }
}

func TestCreateFilePath(t *testing.T) {
	path := image.CreateFilePath("test_img", "png")
	if path != "upload/images/test_img.png" {
		t.Fatal("Incorrect path created")
	}
}

func TestValidateFile(t *testing.T) {
	err := image.ValidateFile("file.png")
	if err != nil {
		t.Fatal("Failed to validate a valid filename")
	}

	err = image.ValidateFile("file.jpg")
	if err != nil {
		t.Fatal("Failed to validate a valid filename")
	}

	err = image.ValidateFile("file_name.jpeg")
	if err != nil {
		t.Fatal("Failed to validate a valid filename")
	}
}

func TestValidateFileWithInvalidFileName(t *testing.T) {
	err := image.ValidateFile("file.pdf")
	if err == nil {
		t.Fatal("Failed to validate a valid filename")
	}

	err = image.ValidateFile(";file.png")
	if err == nil {
		t.Fatal("Failed to validate a valid filename")
	}

	err = image.ValidateFile("file.tiff")
	if err == nil {
		t.Fatal("Failed to validate a valid filename")
	}
}

func TestGetFileExtension(t *testing.T) {
	ext := image.GetFileExtension("file.png")
	if ext != "png" {
		t.Fatal("Extension not returned")
	}
}

func TestGetFileExtensionWithNoExtension(t *testing.T) {
	ext := image.GetFileExtension("file")
	if ext != "" {
		t.Fatal("Failed to return empty string for missing extension")
	}
}
