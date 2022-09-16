package mock

import "github.com/v24Zer0/media-share/post-service/post"

var posts []post.Post = []post.Post{
	{
		ID:        "",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "",
		ImageID:   "",
	},
	{
		ID:        "",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "",
		ImageID:   "",
	},
	{
		ID:        "",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "",
		ImageID:   "",
	},
	{
		ID:        "",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "",
		ImageID:   "",
	},
	{
		ID:        "",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "",
		ImageID:   "",
	},
}

func GetMockPosts() []post.Post {
	return posts
}
