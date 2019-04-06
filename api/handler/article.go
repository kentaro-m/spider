package handler

import (
	"net/http"
	"github.com/go-chi/render"
	"github.com/kentaro-m/spider/api/driver"
	repository "github.com/kentaro-m/spider/api/repository"
	article "github.com/kentaro-m/spider/api/repository/article"
)

func NewArticleHandler(db *driver.DB) *Article {
	return &Article{
		repo: article.NewMySQLArticleRepository(db.SQL),
	}
}

type Article struct {
	repo repository.ArticleRepository
}

func (a *Article) Get(w http.ResponseWriter, r *http.Request)  {
	payload, _ := a.repo.Get(r.Context())

	render.JSON(w, r, payload)
}

func (a *Article) Create(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created an Article successfully"
	render.JSON(w, r, response)
}
