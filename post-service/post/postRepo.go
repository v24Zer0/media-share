package post

import "gorm.io/gorm"

type PostRepo struct {
	db *gorm.DB
}

func (r PostRepo) GetPosts() {

}

func (r PostRepo) CreatePost(post *Post) {

}

func (r PostRepo) DeletePost(id string) {

}

func NewPostRepo() *PostRepo {
	return &PostRepo{}
}
