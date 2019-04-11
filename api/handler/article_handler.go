package handler

import (
	"encoding/json"
	"github.com/kentaro-m/spider/api/model"
	"log"
	"net/http"
)

func NewArticleHandler(m model.ArticleModel) ArticleHandler {
	return &articleHandler{
		model: m,
	}
}

type ArticleHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type articleHandler struct {
	model model.ArticleModel
}

// GetArticle godoc
// @Summary Get articles
// @Tags articles
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Article
// @Router /articles [get]
func (a *articleHandler) Get(w http.ResponseWriter, r *http.Request) {
	payload, err := a.model.Get(r.Context())

	if err != nil {
		log.Printf("Error: %+v\n", err)
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondWithJSON(w, http.StatusOK, payload)
}

// CreateArticle godoc
// @Summary Add a new article
// @Tags articles
// @Accept  json
// @Produce  json
// @Param   article body entity.Article true  "article"
// @Success 200
// @Router /articles [post]
func (a *articleHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := a.model.Create(r.Context(), r)

	if err != nil {
		log.Printf("Error: %+v\n", err)
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}
