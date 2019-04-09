package model

import (
	"context"
	"encoding/json"
	"github.com/kentaro-m/spider/api/entity"
	"github.com/kentaro-m/spider/api/repository"
	"github.com/satori/go.uuid"
	"golang.org/x/xerrors"
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

	if err != nil {
		return nil, xerrors.Errorf("failed to get articles from DB: %w", err)
	}

	return payload, err
}

func (a articleModel) Create(ctx context.Context, r *http.Request) error {
	timeStamp := time.Now().UTC().In(time.FixedZone("Asia/Tokyo", 9*60*60))

	article := entity.Article{
		ID:        uuid.NewV4().String(),
		CreatedAt: timeStamp,
		UpdatedAt: timeStamp,
	}

	err := json.NewDecoder(r.Body).Decode(&article)

	if err != nil {
		return xerrors.Errorf("failed to read JSON-encoded value: %w", err)
	}

	err = a.repo.Create(r.Context(), &article)

	if err != nil {
		return xerrors.Errorf("failed to insert an article to DB: %w", err)
	}

	return err
}
