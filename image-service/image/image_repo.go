package image

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// hold db connection
type ImageRepo struct {
	db *gorm.DB
}

func ConstructDBPath() string {
	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	return fmt.Sprintf("%s:%s@tcp%s/%s", dbUser, dbPass, dbURL, dbName)
}

func NewImageRepo() *ImageRepo {
	dbPath := ConstructDBPath()

	db, err := gorm.Open(mysql.Open(dbPath))
	if err != nil {
		log.Panicln("Failed to connect to database")
	}

	return &ImageRepo{
		db: db,
	}
}

func (repo ImageRepo) GetImage(postID string) (string, error) {
	repo.db.First(&Image{PostID: postID})
	return "", nil
}

func (repo ImageRepo) CreateImage(image *Image) (*Image, error) {
	repo.db.Create(image)
	return image, nil
}

func (repo ImageRepo) DeleteImage(postID string) (*Image, error) {
	repo.db.Delete(&Image{PostID: postID})
	return &Image{}, nil
}
