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

	userID = "0003"
	posts, _ = service.GetPosts(userID)
	if len(*posts) != 0 {
		t.Fatalf("Posts returned for unknown user: %s", userID)
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
		ID:        "",
		Title:     "",
		CreatedAt: "",
		UserID:    "0002",
		CreatedBy: "user2",
	}

	// Should fail to create post and return err with message "missing field"
	err = service.CreatePost(&post2)
	if err == nil {
		t.Fatal("No error when creating post with missing field")
	}
}

func TestServiceDeletePost(t *testing.T) {
	repo := mock.NewMockRepo()
	idProvider := mock.MockIDProvider{}

	service := post.NewPostService(repo, idProvider)

	post1 := post.Post{
		ID:        "1",
		Title:     "Post title 1",
		CreatedAt: "",
		UserID:    "0001",
		CreatedBy: "user1",
	}

	// should delete post with valid id
	err := service.DeletePost(&post1)
	if err != nil {
		t.Fatalf("Failed to delete post with id: %s", post1.ID)
	}

	post2 := post.Post{
		ID:        "1",
		Title:     "Post title 6",
		CreatedAt: "",
		UserID:    "0002",
		CreatedBy: "user2",
	}

	// should fail to delete post with unknown id
	err = service.DeletePost(&post2)
	if err == nil {
		t.Fatalf("No error when deleting post not belonging to user %s with id: %s", post2.UserID, post2.ID)
	}

	post3 := post.Post{
		ID:        "",
		Title:     "Post title 6",
		CreatedAt: "",
		UserID:    "0002",
		CreatedBy: "user2",
	}

	// should fail to delete post with missing ID
	err = service.DeletePost(&post3)
	if err == nil {
		t.Fatal("No error when deleting post with missing field")
	}
}
