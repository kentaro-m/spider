package article

import (
	"github.com/mholt/binding"
	"golang.org/x/xerrors"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"time"
)

type getArticleForm struct {
	Since time.Time `example:"2019-01-19T14:13:01Z"`
	Until time.Time `example:"2019-01-19T14:13:01Z"`
	Limit int `validate:"omitempty,min=1,max=50" example:"50"`
	Sort string `validate:"omitempty,oneof=desc asc" example:"desc"`
}

type GetArticleForm struct {
	getArticleForm
}

func (g *GetArticleForm) Since() time.Time {
	return g.getArticleForm.Since
}

func (g *GetArticleForm) Until() time.Time {
	return g.getArticleForm.Until
}

func (g *GetArticleForm) Limit() int {
	if g.getArticleForm.Limit == 0 {
		g.getArticleForm.Limit = 50
	}

	return g.getArticleForm.Limit
}

func (g *GetArticleForm) Sort() string {
	if g.getArticleForm.Sort == "" {
		g.getArticleForm.Sort = "desc"
	}

	return g.getArticleForm.Sort
}

func (g *GetArticleForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&g.getArticleForm.Since: binding.Field{
			Form: "since",
			Required: false,
		},
		&g.getArticleForm.Until: binding.Field{
			Form: "until",
			Required: false,
		},
		&g.getArticleForm.Limit: binding.Field{
			Form: "limit",
			Required: false,
		},
		&g.getArticleForm.Sort: binding.Field{
			Form: "sort",
			Required: false,
		},
	}
}

const (
	ErrSinceValidationFailed = "since should be a ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)"
	ErrUntilValidationFailed = "until should be a ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)"
	ErrLimitValidationFailed = "limit should be a number (min 1 and max 50)"
	ErrSortValidationFailed = "sort can be a one of asc or desc"
)

func (g *GetArticleForm) Validate(r *http.Request) (string, error) {
	errs := binding.URL(r, g)

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
			return msg, xerrors.Errorf("failed to bind request params: %w", e)
		}
	}

	validate := validator.New()
	err := validate.Struct(g)

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

			return msg, xerrors.Errorf("failed to validate request params: %w", err)
		}
	}

	return "", nil
}

type createArticleForm struct {
	Title     string    `json:"title" validate:"required" example:"AWS CDKでサーバーレスアプリケーションのデプロイを試す"`
	URL       string    `json:"url" validate:"required,url" example:"https://blog.kentarom.com/learn-aws-cdk/"`
	PubDate   time.Time `json:"pub_date" validate:"required" example:"2019-01-19T14:13:01Z"`
	SiteTitle string    `json:"site_title" example:"Learn Something New"`
	SiteURL   string    `json:"site_url" example:"https://blog.kentarom.com"`
}

type CreateArticleForm struct {
	createArticleForm
}

func (c *CreateArticleForm) Title() string {
	return c.createArticleForm.Title
}

func (c *CreateArticleForm) URL() string {
	return c.createArticleForm.URL
}

func (c *CreateArticleForm) PubDate() time.Time {
	return c.createArticleForm.PubDate
}

func (c *CreateArticleForm) SiteTitle() string {
	return c.createArticleForm.SiteTitle
}

func (c *CreateArticleForm) SiteURL() string {
	return c.createArticleForm.SiteURL
}

func (c *CreateArticleForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&c.createArticleForm.Title: binding.Field{
			Form: "title",
			Required: true,
		},
		&c.createArticleForm.URL: binding.Field{
			Form: "url",
			Required: true,
		},
		&c.createArticleForm.PubDate: binding.Field{
			Form: "pub_date",
			Required: true,
		},
		&c.createArticleForm.SiteTitle: binding.Field{
			Form: "site_title",
			Required: false,
		},
		&c.createArticleForm.SiteURL: binding.Field{
			Form: "site_url",
			Required: false,
		},
	}
}

const (
	ErrTitleValidationFailed = "title should be a string"
	ErrURLValidationFailed = "url should be a url format (http://)"
	ErrPubDateValidationFailed = "pub_date should be a ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)"
	ErrSiteTitleValidationFailed = "site_title should be a string"
	ErrSiteURLValidationFailed = "site_url should be a url format (http://)"
)

func (c *CreateArticleForm) Validate(r *http.Request) (string, error) {
	errs := binding.Json(r, c)

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
				case "site_title":
					msg = ErrSiteTitleValidationFailed
				case "site_url":
					msg = ErrSiteURLValidationFailed
				}
			}
			return msg, xerrors.Errorf("failed to bind request params: %w", e)
		}
	}

	validate := validator.New()
	err := validate.Struct(c)

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
			case "SiteTitle":
				msg = ErrSiteTitleValidationFailed
			case "SiteURL":
				msg = ErrSiteURLValidationFailed
			}
		}

		return msg, xerrors.Errorf("failed to validate request params: %w", err)
	}

	return "", nil
}
