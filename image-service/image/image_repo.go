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
	image := &Image{PostID: postID}
	res := repo.db.First(image)
	return image.Path, res.Error
}

func (repo ImageRepo) CreateImage(image *Image) error {
	res := repo.db.Create(image)
	return res.Error
}

func (repo ImageRepo) DeleteImage(postID string) error {
	res := repo.db.Delete(&Image{PostID: postID})
	return res.Error
}
