package post_test

import (
	"testing"

	"github.com/v24Zer0/media-share/post-service/mock"
	"github.com/v24Zer0/media-share/post-service/post"
)

func TestRepoGetPosts(t *testing.T) {

}

func TestRepoCreatePost(t *testing.T) {
	repo := mock.NewMockRepo()
	post := post.Post{
		ID:        "6",
		Title:     "Title 1",
		CreatedAt: "",
		UserID:    "0001",
		CreatedBy: "user1",
	}

	// Valid CreatePost()
	err := repo.CreatePost(&post)
	if err != nil {
		t.Fatal("Failed to create valid post")
	}

	// post.ID = "1"

	// Invalid CreatePost()
	// err = repo.CreatePost(&post)
}

func TestRepoDeletePost(t *testing.T) {

}
