package mock

import "github.com/v24Zer0/media-share/post-service/post"

var posts []post.Post = []post.Post{
	{
		ID:        "1",
		Title:     "Title3",
		CreatedAt: "",
		CreatedBy: "User1",
		ImageID:   "img_id1",
	},
	{
		ID:        "2",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "User2",
		ImageID:   "img_id2",
	},
	{
		ID:        "3",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "User3",
		ImageID:   "img_id3",
	},
	{
		ID:        "4",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "User4",
		ImageID:   "img_id4",
	},
	{
		ID:        "5",
		Title:     "",
		CreatedAt: "",
		CreatedBy: "User5",
		ImageID:   "img_id5",
	},
}

func GetMockPosts() *[]post.Post {
	return &posts
}
