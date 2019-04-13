package form

import (
	"github.com/mholt/binding"
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
