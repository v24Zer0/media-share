package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/v24Zer0/media-share/post-service/post"
	"github.com/v24Zer0/media-share/post-service/router"
	"github.com/v24Zer0/media-share/post-service/util"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load .env file")
	}

	repo := post.NewPostRepo()
	service := post.NewPostService(repo, util.DefaultIDProvider{})
	handler := post.NewPostHandler(service)

	r := router.NewRouter(handler)
	r.Run(":8080")
}
