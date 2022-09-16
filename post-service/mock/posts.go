package mock

import "github.com/v24Zer0/media-share/post-service/post"

var posts []post.Post = []post.Post{
	{
		ID:        "1",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "",
		ImageID:   "",
	},
	{
		ID:        "2",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "",
		ImageID:   "",
	},
	{
		ID:        "3",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "",
		ImageID:   "",
	},
	{
		ID:        "4",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "",
		ImageID:   "",
	},
	{
		ID:        "5",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "",
		ImageID:   "",
	},
}

func GetMockPosts() []post.Post {
	return posts
}
