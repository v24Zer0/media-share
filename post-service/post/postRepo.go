package post

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PostRepo struct {
	db *gorm.DB
}

func ConstructDBPath() string {
	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	return fmt.Sprintf("%s:%s@tcp%s/%s", dbUser, dbPass, dbURL, dbName)
}

func NewPostRepo() *PostRepo {
	dbPath := ConstructDBPath()

	db, err := gorm.Open(mysql.Open(dbPath))
	if err != nil {
		log.Panicln("Failed to connect to database")
	}

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

func (repo PostRepo) DeletePost(post *Post) error {
	res := repo.db.Delete(post)
	return res.Error
}
