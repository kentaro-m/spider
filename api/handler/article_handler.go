package handler

import (
	"encoding/json"
	"net/http"
	"github.com/kentaro-m/spider/api/driver"
	"github.com/kentaro-m/spider/api/repository"
	"github.com/kentaro-m/spider/api/entity"
	"github.com/satori/go.uuid"
	"time"
)

func NewArticleHandler(db *driver.DB) *Article {
	return &Article{
		repo: repository.NewArticleRepository(db.SQL),
	}
}

type Article struct {
	repo repository.ArticleRepository
}

func (a *Article) Get(w http.ResponseWriter, r *http.Request)  {
	payload, _ := a.repo.Get(r.Context())
	respondwithJSON(w, http.StatusOK, payload)
}

func (a *Article) Create(w http.ResponseWriter, r *http.Request) {
	timeStamp := time.Now().UTC().In(time.FixedZone("Asia/Tokyo", 9*60*60))

	article := entity.Article{
		ID: uuid.NewV4().String(),
		CreatedAt: timeStamp,
		UpdatedAt: timeStamp,
	}
	json.NewDecoder(r.Body).Decode(&article)
	err := a.repo.Create(r.Context(), &article)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
