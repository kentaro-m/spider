package form

import (
	"github.com/mholt/binding"
	"golang.org/x/xerrors"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"time"
)

type GetArticleForm struct {
	Since time.Time
	Until time.Time
	Limit int `validate:"omitempty,min=1,max=50"`
	Sort string `validate:"omitempty,oneof=desc asc"`
}

func (g *GetArticleForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&g.Since: binding.Field{
			Form: "since",
			Required: false,
		},
		&g.Until: binding.Field{
			Form: "until",
			Required: false,
		},
		&g.Limit: binding.Field{
			Form: "limit",
			Required: false,
		},
		&g.Sort: binding.Field{
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

type CreateArticleForm struct {
	Title     string    `json:"title" validate:"required" example:"AWS CDKでサーバーレスアプリケーションのデプロイを試す"`
	URL       string    `json:"url" validate:"required,url" example:"https://blog.kentarom.com/learn-aws-cdk/"`
	PubDate   time.Time `json:"pub_date" validate:"required" example:"2019-01-19T14:13:01Z"`
}

func (c *CreateArticleForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&c.Title: binding.Field{
			Form: "title",
			Required: true,
		},
		&c.URL: binding.Field{
			Form: "url",
			Required: true,
		},
		&c.PubDate: binding.Field{
			Form: "pub_date",
			Required: true,
		},
	}
}

const (
	ErrTitleValidationFailed = "title should be a string"
	ErrURLValidationFailed = "url should be a url format (http://)"
	ErrPubDateValidationFailed = "pub_date should be a ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)"
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
			}
		}

		return msg, xerrors.Errorf("failed to validate request params: %w", err)
	}

	return "", nil
}
