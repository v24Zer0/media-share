package post

type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	ImageID   string `json:"imageID"`
}
