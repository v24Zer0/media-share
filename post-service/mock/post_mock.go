package mock

import (
	"github.com/v24Zer0/media-share/post-service/post"
	"github.com/v24Zer0/media-share/post-service/user"
)

var posts []post.Post = []post.Post{
	{
		ID:        "1",
		Title:     "Post title 1",
		CreatedAt: "",
		CreatedBy: user.User{
			ID:       "0001",
			Username: "user1",
		},
		ImageID: "img_id1",
	},
	{
		ID:        "2",
		Title:     "Post title 2",
		CreatedAt: "",
		CreatedBy: user.User{
			ID:       "0001",
			Username: "user1",
		},
		ImageID: "img_id2",
	},
	{
		ID:        "3",
		Title:     "Post title 3",
		CreatedAt: "",
		CreatedBy: user.User{
			ID:       "0001",
			Username: "user1",
		},
		ImageID: "img_id3",
	},
	{
		ID:        "4",
		Title:     "Post title 4",
		CreatedAt: "",
		CreatedBy: user.User{
			ID:       "0002",
			Username: "user2",
		},
		ImageID: "img_id4",
	},
	{
		ID:        "5",
		Title:     "Post title 5",
		CreatedAt: "",
		CreatedBy: user.User{
			ID:       "0002",
			Username: "user2",
		},
		ImageID: "img_id5",
	},
}

func GetMockPosts() *[]post.Post {
	return &posts
}
