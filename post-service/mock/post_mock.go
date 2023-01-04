package mock

import (
	"github.com/v24Zer0/media-share/post-service/post"
)

var posts []post.Post = []post.Post{
	{
		ID:        "1",
		Title:     "Post title 1",
		CreatedAt: "",
		UserID:    "0001",
		CreatedBy: "user1",
	},
	{
		ID:        "2",
		Title:     "Post title 2",
		CreatedAt: "",
		UserID:    "0001",
		CreatedBy: "user1",
	},
	{
		ID:        "3",
		Title:     "Post title 3",
		CreatedAt: "",
		UserID:    "0001",
		CreatedBy: "user1",
	},
	{
		ID:        "4",
		Title:     "Post title 4",
		CreatedAt: "",
		UserID:    "0002",
		CreatedBy: "user2",
	},
	{
		ID:        "5",
		Title:     "Post title 5",
		CreatedAt: "",
		UserID:    "0002",
		CreatedBy: "user2",
	},
}

func GetMockPosts() *[]post.Post {
	return &posts
}
