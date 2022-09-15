package post

type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedOn string `json:"createdOn"`
	CreatedBy string `json:"createdBy"`
	ImageID   string `json:"imageID"`
}
