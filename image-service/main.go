package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/v24Zer0/media-share/image-service/image"
	"github.com/v24Zer0/media-share/image-service/router"
	"github.com/v24Zer0/media-share/image-service/util"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading environment variables")
	}

	repo := image.NewImageRepo()
	service := image.NewImageService(repo, util.DefaultProvider{})
	handler := image.NewImageHandler(service)

	r := router.NewRouter(handler)
	r.Run(":8081")
}
