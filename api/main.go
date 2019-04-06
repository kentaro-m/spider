package main

import (
	"github.com/kentaro-m/spider/api/handler"

	"net/http"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/articles", handler.GetArticles)
	r.Post("/articles", handler.CreateArticle)

	http.ListenAndServe(":8080", r)
}
