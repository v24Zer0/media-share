package post

type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	UserID    string `json:"userId"`
	CreatedBy string `json:"createdBy"`
	ImageID   string `json:"imageID"`
}
