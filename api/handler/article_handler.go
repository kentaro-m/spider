package handler

import (
	"encoding/json"
	"github.com/kentaro-m/spider/api/model"
	"github.com/kentaro-m/spider/api/form"
	"golang.org/x/xerrors"
	"log"
	"net/http"
	"gopkg.in/go-playground/validator.v9"
	"github.com/mholt/binding"
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

const (
	ErrSinceValidationFailed = "since should be a ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)"
	ErrUntilValidationFailed = "until should be a ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)"
	ErrLimitValidationFailed = "limit should be a number (min 1 and max 50)"
	ErrSortValidationFailed = "sort can be a one of asc or desc"
	ErrTitleValidationFailed = "title should be a string"
	ErrURLValidationFailed = "url should be a url format (http://)"
	ErrPubDateValidationFailed = "pub_date should be a ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)"
)

// GetArticle godoc
// @Summary Get articles
// @Tags articles
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Article
// @Router /articles [get]
func (a *articleHandler) Get(w http.ResponseWriter, r *http.Request) {
	getArticleForm := new(form.GetArticleForm)
	errs := binding.URL(r, getArticleForm)

	if errs != nil {
		var msg string

		for _, e := range errs.(binding.Errors) {
			for _, fieldName := range e.Fields() {
				switch fieldName {
				case "since":
					msg = ErrSinceValidationFailed
				case "until":
					msg = ErrUntilValidationFailed
				case "limit":
					msg = ErrLimitValidationFailed
				case "sort":
					msg = ErrSortValidationFailed
				}
			}
			log.Printf("Error: %+v\n", xerrors.Errorf("failed to bind request params: %w", e))
			respondWithError(w, http.StatusBadRequest, msg)
			return
		}
	}

	validate := validator.New()
	err := validate.Struct(getArticleForm)

	if err != nil {
		var msg string
		for _, e := range err.(validator.ValidationErrors) {
			fieldName := e.Field()
			switch fieldName {
			case "Since":
				msg = ErrSinceValidationFailed
			case "Until":
				msg = ErrUntilValidationFailed
			case "Limit":
				msg = ErrLimitValidationFailed
			case "Sort":
				msg = ErrSortValidationFailed
			}
		}

		log.Printf("Error: %+v\n", xerrors.Errorf("failed to validate request params: %w", err))
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
// @Param   article body entity.Article true  "article"
// @Success 200
// @Router /articles [post]
func (a *articleHandler) Create(w http.ResponseWriter, r *http.Request) {
	createArticleForm := new(form.CreateArticleForm)
	errs := binding.Json(r, createArticleForm)

	if errs != nil {
		var msg string
		for _, e := range errs.(binding.Errors) {
			for _, fieldName := range e.Fields() {
				switch fieldName {
				case "title":
					msg = ErrTitleValidationFailed
				case "url":
					msg = ErrURLValidationFailed
				case "pub_date":
					msg = ErrPubDateValidationFailed
				}
			}
			log.Printf("Error: %+v\n", xerrors.Errorf("failed to bind request params: %w", e))
			respondWithError(w, http.StatusBadRequest, msg)
			return
		}
	}

	validate := validator.New()
	err := validate.Struct(createArticleForm)

	if err != nil {
		var msg string
		for _, e := range err.(validator.ValidationErrors) {
			fieldName := e.Field()
			switch fieldName {
			case "Title":
				msg = ErrTitleValidationFailed
			case "URL":
				msg = ErrURLValidationFailed
			case "PubDate":
				msg = ErrPubDateValidationFailed
			}
		}

		log.Printf("Error: %+v\n", xerrors.Errorf("failed to validate request params: %w", err))
		respondWithError(w, http.StatusBadRequest, msg)
		return
	}

	err = a.model.Create(r.Context(), r, createArticleForm)

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
