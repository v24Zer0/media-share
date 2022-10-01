package post_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/v24Zer0/media-share/post-service/mock"
	"github.com/v24Zer0/media-share/post-service/post"
	"github.com/v24Zer0/media-share/post-service/router"
)

func TestNewPostHandler(t *testing.T) {

}

func TestHandlerGetPosts(t *testing.T) {
	service := mock.NewMockService(mock.NewMockRepo(), mock.MockIDProvider{})
	handler := post.NewPostHandler(service)

	router := router.NewRouter(handler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/post/0001", nil)
	router.ServeHTTP(w, req)

	posts := []post.Post{}
	json.Unmarshal(w.Body.Bytes(), &posts)

	for _, p := range posts {
		if p.UserID != "0001" {
			t.Fatal("Incorrect post retrieved")
		}
	}
}

func TestHandlerCreatePost(t *testing.T) {
	service := mock.NewMockService(mock.NewMockRepo(), mock.MockIDProvider{})
	handler := post.NewPostHandler(service)

	router := router.NewRouter(handler)

	post1 := post.Post{
		ID:        "",
		Title:     "Post title 6",
		CreatedAt: "",
		UserID:    "0002",
		CreatedBy: "user2",
	}

	b, _ := json.Marshal(&post1)
	body := bytes.NewReader(b)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/post", body)

	router.ServeHTTP(w, req)

	// should create post and return code 201
	if w.Result().StatusCode != 201 {
		t.Fatal("Failed  to create post with valid fields")
	}

	post2 := post.Post{
		ID:        "",
		Title:     "",
		CreatedAt: "",
		UserID:    "0002",
		CreatedBy: "user2",
	}

	b, _ = json.Marshal(&post2)
	body = bytes.NewReader(b)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/post", body)

	router.ServeHTTP(w, req)

	// should create post and return code 201
	if w.Result().StatusCode != 400 {
		t.Fatal("Created post with missing fields")
	}
}

func TestHandlerDeletePost(t *testing.T) {
	service := mock.NewMockService(mock.NewMockRepo(), mock.MockIDProvider{})
	handler := post.NewPostHandler(service)

	router := router.NewRouter(handler)

	post1 := post.Post{
		ID:        "1",
		Title:     "Post title 1",
		CreatedAt: "",
		UserID:    "0001",
		CreatedBy: "user1",
	}

	b, _ := json.Marshal(&post1)
	body := bytes.NewReader(b)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/post", body)

	router.ServeHTTP(w, req)

	// should delete post with valid ID and UserID
	if w.Result().StatusCode != 200 {
		t.Fatalf("Failed to delete post with ID: %s", post1.ID)
	}

	post2 := post.Post{
		ID:        "1",
		Title:     "Post title 1",
		CreatedAt: "",
		UserID:    "0002",
		CreatedBy: "user1",
	}

	b, _ = json.Marshal(&post2)
	body = bytes.NewReader(b)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/post", body)

	router.ServeHTTP(w, req)

	// should fail to delete post because Post does not belong to the user
	if w.Result().StatusCode != 400 {
		t.Fatalf("Deleted post not belonging to User: %s with ID: %s", post2.UserID, post2.ID)
	}

	post3 := post.Post{
		ID:        "",
		Title:     "Post title 1",
		CreatedAt: "",
		UserID:    "0001",
		CreatedBy: "user1",
	}

	b, _ = json.Marshal(&post3)
	body = bytes.NewReader(b)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/post", body)

	router.ServeHTTP(w, req)

	// should fail to delete post because a field is missing
	if w.Result().StatusCode != 400 {
		t.Fatal("Deleted post with missing fields")
	}
}
