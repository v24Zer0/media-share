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

func (repo PostRepo) GetPosts() (*[]Post, error) {
	posts := []Post{}
	res := repo.db.Find(&posts, Post{ID: ""})
	if res.Error != nil {
		return &[]Post{}, res.Error
	}
	return &posts, nil
}

func (repo PostRepo) CreatePost(post *Post) error {
	res := repo.db.Create(post)
	return res.Error
}

func (repo PostRepo) DeletePost(id string) error {
	res := repo.db.Delete(&Post{ID: id})
	return res.Error
}
