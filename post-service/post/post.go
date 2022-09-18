package post

import "github.com/v24Zer0/media-share/post-service/user"

type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt string    `json:"createdAt"`
	CreatedBy user.User `json:"createdBy"`
	ImageID   string    `json:"imageID"`
}
