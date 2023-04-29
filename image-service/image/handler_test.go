package image_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/v24Zer0/media-share/image-service/image"
	"github.com/v24Zer0/media-share/image-service/mock"
	"github.com/v24Zer0/media-share/image-service/response"
	"github.com/v24Zer0/media-share/image-service/router"
	"github.com/v24Zer0/media-share/image-service/util"
)

func TestHandler_GetImage(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &util.DefaultProvider{})
	handler := image.NewImageHandler(service)
	router := router.NewRouter(handler)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/image/post_03", nil)
	router.ServeHTTP(recorder, req)

	if recorder.Result().StatusCode != 200 {
		t.Fatal("Request failed")
	}

	body := recorder.Result().Body
	defer body.Close()

	// Convert body to bytes to compare to actual image
	respImg, err := ioutil.ReadAll(body)
	if err != nil {
		t.Fatal("Failed to read bytes from image")
	}

	// Fetch actual Image bytes to compare
	actaulImg, err := ioutil.ReadFile("upload/images/02f0d331-3ed1-43cc-8f88-0eb4b9501d8c.jpeg")
	if err != nil {
		t.Fatal("Failed to read actual image")
	}

	if !bytes.Equal(respImg, actaulImg) {
		t.Fatal("Response image does not match actual image")
	}
}

func TestHandler_GetImageWithMissingPostID(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &util.DefaultProvider{})
	handler := image.NewImageHandler(service)
	router := router.NewRouter(handler)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/image", nil)
	router.ServeHTTP(recorder, req)

	if recorder.Result().StatusCode != 404 {
		t.Fatal("Error not returned for invalid route")
	}
}

func TestHandler_GetImageWithErrorFromServiceCall(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &util.DefaultProvider{})
	handler := image.NewImageHandler(service)
	router := router.NewRouter(handler)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/image/post_04", nil)
	router.ServeHTTP(recorder, req)

	// Should be bad request for invalid postID
	if recorder.Result().StatusCode != 400 {
		t.Fatal("Bad Request not returned for invalid postID")
	}
}

func TestHandler_CreateImage(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &mock.MockIDProvider{})
	handler := image.NewImageHandler(service)
	router := router.NewRouter(handler)

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)

	err := w.WriteField("postID", "post_01")
	if err != nil {
		t.Fatal("failed to write form field")
	}

	fileWriter, err := w.CreateFormFile("image", "image.jpg")
	if err != nil {
		t.Fatal("Failed to create form file")
	}

	_, err = fileWriter.Write(mockImage)
	if err != nil {
		t.Fatal("Failed to write image bytes")
	}

	w.Close()

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/image", body)
	req.Header.Add("Content-Type", w.FormDataContentType())
	router.ServeHTTP(recorder, req)

	if recorder.Result().StatusCode != 201 {
		t.Fatal("Failed to return status 201 for a created image")
	}

	RestoreStateAfterCreate()
}

func TestHandler_CreateImageWithMissingPostID(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &mock.MockIDProvider{})
	handler := image.NewImageHandler(service)
	router := router.NewRouter(handler)

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)

	fileWriter, err := w.CreateFormFile("image", "image.jpg")
	if err != nil {
		t.Fatal("Failed to create form file")
	}

	_, err = fileWriter.Write(mockImage)
	if err != nil {
		t.Fatal("Failed to write image bytes")
	}

	w.Close()

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/image", body)
	req.Header.Add("Content-Type", w.FormDataContentType())
	router.ServeHTTP(recorder, req)

	if recorder.Result().StatusCode != 400 {
		t.Fatal("Request to create image with invalid postID did not fail")
	}

	// Compare error messages
	message := response.ErrorResponse{}
	json.Unmarshal(recorder.Body.Bytes(), &message)
	if message.Message != "missing postID" {
		t.Fatal("Incorrect error message returned for missing postID")
	}
}

func TestHandler_CreateImageWithMissingImage(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &mock.MockIDProvider{})
	handler := image.NewImageHandler(service)
	router := router.NewRouter(handler)

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)

	err := w.WriteField("postID", "post_01")
	if err != nil {
		t.Fatal("failed to write form field")
	}

	w.Close()

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/image", body)
	req.Header.Add("Content-Type", w.FormDataContentType())
	router.ServeHTTP(recorder, req)

	if recorder.Result().StatusCode != 400 {
		t.Fatal("Failed to return status 400 for missing image")
	}

	// Compare error messages
	message := response.ErrorResponse{}
	json.Unmarshal(recorder.Body.Bytes(), &message)
	if message.Message != "missing image" {
		t.Fatal("Incorrect error message returned for missing image")
	}
}

func TestHandler_CreateImageWithErrorReadingImage(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &mock.MockIDProvider{})
	handler := image.NewImageHandler(service)
	router := router.NewRouter(handler)

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)

	w.WriteField("postID", "post_01")
	w.CreateFormFile("image", "image.jpg")

	w.Close()

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/image", body)
	req.Header.Add("Content-Type", w.FormDataContentType())
	router.ServeHTTP(recorder, req)

	if recorder.Result().StatusCode != 400 {
		t.Fatal("Failed to return status 400 for failure to read image")
	}

	// Compare error messages
	message := response.ErrorResponse{}
	json.Unmarshal(recorder.Body.Bytes(), &message)
	if message.Message != "failed to read image" {
		t.Fatal("Incorrect error message returned for failure to read image")
	}
}

func TestHandler_CreateImageWithErrorFromServiceCall(t *testing.T) {
	service := image.NewImageService(mock.NewMockRepo(), &mock.MockIDProvider{})
	handler := image.NewImageHandler(service)
	router := router.NewRouter(handler)

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)

	err := w.WriteField("postID", "post_01")
	if err != nil {
		t.Fatal("failed to write form field")
	}

	fileWriter, err := w.CreateFormFile("image", "image.pdf")
	if err != nil {
		t.Fatal("Failed to create form file")
	}

	_, err = fileWriter.Write(mockImage)
	if err != nil {
		t.Fatal("Failed to write image bytes")
	}

	w.Close()

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/image", body)
	req.Header.Add("Content-Type", w.FormDataContentType())
	router.ServeHTTP(recorder, req)

	if recorder.Result().StatusCode != 400 {
		t.Fatal("Failed to return status 400 for error in service call")
	}

	// Compare error messages
	message := response.ErrorResponse{}
	json.Unmarshal(recorder.Body.Bytes(), &message)
	if message.Message != "invalid filename format" {
		t.Fatal("Incorrect error message returned for error in service call")
	}
}

// func TestHandler_DeleteImage(t *testing.T) {
// 	service := image.NewImageService(mock.NewMockRepo(), &util.DefaultProvider{})
// 	handler := image.NewImageHandler(service)
// 	router := router.NewRouter(handler)

// 	recorder := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/image/post_03", nil)
// 	router.ServeHTTP(recorder, req)
// }

// func TestHandler_DeleteImageWithInvalidRequestBody(t *testing.T) {
// 	service := image.NewImageService(mock.NewMockRepo(), &util.DefaultProvider{})
// 	handler := image.NewImageHandler(service)
// 	router := router.NewRouter(handler)

// 	recorder := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/image", nil)
// 	router.ServeHTTP(recorder, req)
// }

// func TestHandler_DeleteImageWithErrorFromServiceCall(t *testing.T) {
// 	service := image.NewImageService(mock.NewMockRepo(), &util.DefaultProvider{})
// 	handler := image.NewImageHandler(service)
// 	router := router.NewRouter(handler)

// 	recorder := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/image", nil)
// 	router.ServeHTTP(recorder, req)
// }
