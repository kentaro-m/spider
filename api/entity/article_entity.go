package entity

type Article struct {
	ID string `json:"id"`
	Title string `json:"title"`
	URL string `json:"url"`
	PubDate string `json:"pub_date"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}