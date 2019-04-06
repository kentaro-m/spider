package handler

import (
	"encoding/json"
	"github.com/kentaro-m/spider/api/model"
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
	payload, _ := a.model.Get(r.Context())
	respondwithJSON(w, http.StatusOK, payload)
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
