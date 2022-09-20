package post

import (
	"gorm.io/gorm"
)

type PostRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (repo PostRepo) GetPosts(userID string) (*[]Post, error) {
	posts := []Post{}
	res := repo.db.Find(&posts, Post{UserID: userID})
	return &posts, res.Error
}

func (repo PostRepo) CreatePost(post *Post) error {
	res := repo.db.Create(post)
	return res.Error
}

func (repo PostRepo) DeletePost(id string) error {
	res := repo.db.Delete(&Post{ID: id})
	return res.Error
}
