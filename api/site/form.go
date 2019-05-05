package site

type GetSiteForm struct {
	Title     string    `json:"title" validate:"required" example:"Learn Something New"`
}

type CreateSiteForm struct {
	Title     string    `json:"title" validate:"required" example:"Learn Something New"`
	URL       string    `json:"url" validate:"required,url" example:"https://blog.kentarom.com"`
}