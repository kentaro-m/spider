package handler

import (
	"encoding/json"
	"github.com/kentaro-m/spider/api/model"
	"golang.org/x/xerrors"
	"log"
	"net/http"
	"gopkg.in/go-playground/validator.v9"
	"github.com/mholt/binding"
	"time"
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

type GetArticleForm struct {
	Since time.Time
	Until time.Time
	Limit int `validate:"min=1,max=50"`
	Sort string `validate:"oneof=desc asc"`
}

func (g *GetArticleForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&g.Since: "since",
		&g.Until: "until",
		&g.Limit: "limit",
		&g.Sort: "sort",
	}
}

type CreateArticleForm struct {
	Title     string    `json:"title" example:"AWS CDKでサーバーレスアプリケーションのデプロイを試す"`
	URL       string    `json:"url" example:"https://blog.kentarom.com/learn-aws-cdk/"`
	PubDate   time.Time `json:"pub_date" example:"2019-01-19T14:13:01Z"`
}

func (c *CreateArticleForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&c.Title: "title",
		&c.URL: "url",
		&c.PubDate: "pub_date",
	}
}

// GetArticle godoc
// @Summary Get articles
// @Tags articles
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Article
// @Router /articles [get]
func (a *articleHandler) Get(w http.ResponseWriter, r *http.Request) {
	getArticleForm := new(GetArticleForm)
	err := binding.URL(r, getArticleForm)

	if err != nil {
		log.Printf("Error: %+v\n", xerrors.Errorf("failed to bind request params: %w", err))
		respondWithError(w, http.StatusBadRequest, "Server Error")
		return
	}

	validate := validator.New()
	err = validate.Struct(getArticleForm)

	if err != nil {
		var msg string
		for _, e := range err.(validator.ValidationErrors) {
			fieldName := e.Field()
			switch fieldName {
			case "Since":
				msg = "since should be a ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)"
			case "Until":
				msg = "until should be a ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)"
			case "Limit":
				msg = "limit should be a number (min 1 and max 50)"
			case "Sort":
				msg = "sort can be a one of asc or desc"
			}
		}

		log.Printf("Error: %+v\n", xerrors.Errorf("failed to validate request params: %w", err))
		respondWithError(w, http.StatusBadRequest, msg)
		return
	}

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
	createArticleForm := new(CreateArticleForm)
	err := binding.Form(r, createArticleForm)

	if err != nil {
		log.Printf("Error: %+v\n", xerrors.Errorf("failed to bind request params: %w", err))
		respondWithError(w, http.StatusBadRequest, "Server Error")
		return
	}

	validate := validator.New()
	err = validate.Struct(createArticleForm)

	if err != nil {
		var msg string
		for _, e := range err.(validator.ValidationErrors) {
			fieldName := e.Field()
			switch fieldName {
			case "Title":
				msg = "title should be a string"
			case "URL":
				msg = "url should be a string"
			case "PubDate":
				msg = "pub_date should be a ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ"
			}
		}

		log.Printf("Error: %+v\n", xerrors.Errorf("failed to validate request params: %w", err))
		respondWithError(w, http.StatusBadRequest, msg)
		return
	}

	err = a.model.Create(r.Context(), r)

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
