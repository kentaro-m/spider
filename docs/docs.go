// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-04-20 14:26:49.780483 +0900 JST m=+0.357515311

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "info": {
        "contact": {},
        "license": {}
    },
    "paths": {
        "/articles": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Get articles",
                "parameters": [
                    {
                        "type": "string",
                        "default": "2019-01-19T14:13:01Z",
                        "description": "Only articles published at or after this time are returned.",
                        "name": "since",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Only articles published at or before this time are returned.",
                        "name": "until",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "desc",
                            "asc"
                        ],
                        "type": "string",
                        "default": "desc",
                        "description": "The direction of the sort by pub_date",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "maximum": 50,
                        "minimum": 1,
                        "type": "integer",
                        "default": 50,
                        "description": "The number of articles that you can get the result",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/Article"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Add a new article",
                "parameters": [
                    {
                        "description": "article",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/CreateArticleForm"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
