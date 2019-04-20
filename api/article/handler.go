package article

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"log"
	"net/http"
)

func NewArticleHandler(m ArticleModel) ArticleHandler {
	return &articleHandler{
		model: m,
	}
}

type ArticleHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type articleHandler struct {
	model ArticleModel
}

// GetArticle godoc
// @Summary Get articles
// @Tags articles
// @Accept  json
// @Produce  json
// @Param since query string false "Only articles published at or after this time are returned." default(2019-01-19T14:13:01Z)
// @Param until query string false "Only articles published at or before this time are returned."
// @Param sort query string false "The direction of the sort by pub_date" default(desc) Enums(desc, asc)
// @Param limit query int false "The number of articles that you can get the result" default(50) mininum(1) maxinum(50)
// @Success 200 {object} Article
// @Router /articles [get]
func (a *articleHandler) Get(w http.ResponseWriter, r *http.Request) {
	getArticleForm := new(GetArticleForm)
	msg, err := getArticleForm.Validate(r)

	if err != nil {
		log.Printf("Error: %+v\n", xerrors.Errorf("failed to validate form value: %w", err))
		respondWithError(w, http.StatusBadRequest, msg)
		return
	}

	payload, err := a.model.Get(r.Context(), getArticleForm)

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
// @Param   article body CreateArticleForm true  "article"
// @Success 200
// @Router /articles [post]
func (a *articleHandler) Create(w http.ResponseWriter, r *http.Request) {
	createArticleForm := new(CreateArticleForm)
	msg, err := createArticleForm.Validate(r)

	if err != nil {
		log.Printf("Error: %+v\n", xerrors.Errorf("failed to validate form value: %w", err))
		respondWithError(w, http.StatusBadRequest, msg)
		return
	}

	err = a.model.Create(r.Context(), r, createArticleForm)

	if err != nil {
		log.Printf("Error: %+v\n", err)
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Successfully Created"})
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
