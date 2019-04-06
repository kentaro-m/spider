package model

import (
	"context"
	"encoding/json"
	"github.com/kentaro-m/spider/api/entity"
	"github.com/kentaro-m/spider/api/repository"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

type ArticleModel interface {
	Get(ctx context.Context) ([]*entity.Article, error)
	Create(ctx context.Context, r *http.Request) error
}

func NewArticleModel(r repository.ArticleRepository) ArticleModel {
	return &articleModel{
		repo: r,
	}
}

type articleModel struct {
	repo repository.ArticleRepository
}

func (a articleModel) Get(ctx context.Context) ([]*entity.Article, error) {
	payload, err := a.repo.Get(ctx)
	return payload, err
}

func (a articleModel) Create(ctx context.Context, r *http.Request) error {
	timeStamp := time.Now().UTC().In(time.FixedZone("Asia/Tokyo", 9*60*60))

	article := entity.Article{
		ID:        uuid.NewV4().String(),
		CreatedAt: timeStamp,
		UpdatedAt: timeStamp,
	}
	json.NewDecoder(r.Body).Decode(&article)
	err := a.repo.Create(r.Context(), &article)

	return err
}
