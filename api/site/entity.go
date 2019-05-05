package site

import "time"

type Site struct {
	ID        string    `json:"id" example:"faf9c3a7-b3ee-441f-baec-a5b668948382"`
	Title     string    `json:"title" example:"Learn Something New"`
	URL       string    `json:"url" example:"https://blog.kentarom.com"`
	CreatedAt time.Time `json:"created_at" example:"2019-04-06T16:03:31Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-04-06T16:03:31Z"`
}