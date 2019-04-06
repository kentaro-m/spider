package entity

import "time"

type Article struct {
	ID string `json:"id"`
	Title string `json:"title"`
	URL string `json:"url"`
	PubDate string `json:"pub_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}