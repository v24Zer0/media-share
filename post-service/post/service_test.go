package post_test

import (
	"testing"

	"github.com/v24Zer0/media-share/post-service/mock"
	"github.com/v24Zer0/media-share/post-service/post"
)

func TestServiceGetPosts(t *testing.T) {
	repo := mock.NewMockRepo()
	idProvider := mock.MockIDProvider{}

	service := post.NewPostService(repo, idProvider)

	userID := "0001"
	posts, _ := service.GetPosts(userID)
	for _, p := range *posts {
		if p.UserID != userID {
			t.Fatalf("Posts for user(%s) not returned", userID)
		}
	}

	userID = "0002"
	posts, _ = service.GetPosts(userID)
	for _, p := range *posts {
		if p.UserID != userID {
			t.Fatalf("Posts for user(%s) not returned", userID)
		}
	}
}

func TestServiceCreatePost(t *testing.T) {
	repo := mock.NewMockRepo()
	idProvider := mock.MockIDProvider{}

	service := post.NewPostService(repo, idProvider)
	post1 := post.Post{
		ID:        "",
		Title:     "Title 6",
		CreatedAt: "",
		UserID:    "0002",
		CreatedBy: "user2",
	}

	// Should create post and return nil err
	err := service.CreatePost(&post1)
	if err != nil {
		t.Fatalf("Failed to create post with id: %s", post1.ID)
	}

	post2 := post.Post{
		ID:        "1",
		Title:     "Title 6",
		CreatedAt: "",
		UserID:    "0002",
		CreatedBy: "user2",
	}

	// Should fail to create duplicate post
	err = service.CreatePost(&post2)
	if err != nil {
		t.Fatalf("No error when creating duplicate post with id: %s", post2.ID)
	}

	post3 := post.Post{
		ID:        "",
		Title:     "",
		CreatedAt: "",
		UserID:    "0002",
		CreatedBy: "user2",
	}

	// Should fail to create post and return err with message "missing field"
	err = service.CreatePost(&post3)
	if err == nil {
		t.Fatal("No error when creating post with missing field")
	}
}

func TestServiceDeletePost(t *testing.T) {
	repo := mock.NewMockRepo()
	idProvider := mock.MockIDProvider{}

	service := post.NewPostService(repo, idProvider)

	// should delete post with valid id
	id := "1"
	err := service.DeletePost(id)
	if err != nil {
		t.Fatalf("Failed to delete post with id: %s", id)
	}

	// should fail to delete post with unknown id
	id = "6"
	err = service.DeletePost(id)
	if err == nil {
		t.Fatalf("No error when deleting unknown post with id: %s", id)
	}
}
