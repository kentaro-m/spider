package handler

import (
	"net/http"
	"github.com/go-chi/render"
)

type Article struct {
	ID string `json:"id"`
	Title string `json:"title"`
	URL string `json:"url"`
	PubDate string `json:"pub_date"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetArticles(w http.ResponseWriter, r *http.Request)  {
	articles := []Article{
		{
			ID: "10ba038e-48da-487b-96e8-8d3b99b6d18a",
			Title: "go-chiでRest APIを作る",
			URL: "https://blog.kentarom.com/creating-rest-api-with-go-chi/",
			PubDate: "",
			CreatedAt: "",
			UpdatedAt: "",
		},
	}

	render.JSON(w, r, articles)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created an Article successfully"
	render.JSON(w, r, response)
}
