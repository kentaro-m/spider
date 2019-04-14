package model

import (
	"context"
	"github.com/kentaro-m/spider/api/entity"
	"github.com/kentaro-m/spider/api/repository"
	"github.com/kentaro-m/spider/api/form"
	"github.com/satori/go.uuid"
	"golang.org/x/xerrors"
	"net/http"
	"time"
)

type ArticleModel interface {
	Get(ctx context.Context, g *form.GetArticleForm) ([]*entity.Article, error)
	Create(ctx context.Context, r *http.Request, c *form.CreateArticleForm) error
}

func NewArticleModel(r repository.ArticleRepository) ArticleModel {
	return &articleModel{
		repo: r,
	}
}

type articleModel struct {
	repo repository.ArticleRepository
}

func (a articleModel) Get(ctx context.Context, g *form.GetArticleForm) ([]*entity.Article, error) {
	if g.Limit == 0 {
		g.Limit = 50
	}

	if g.Sort == "" {
		g.Sort = "desc"
	}

	if !g.Until.IsZero() && !g.Since.IsZero() {
		payload, err := a.repo.GetArticlesBySinceAndUntil(ctx, g)

		if err != nil {
			return nil, xerrors.Errorf("failed to get articles from DB: %w", err)
		}

		return payload, err
	}

	if g.Until.IsZero() && !g.Since.IsZero() {
		payload, err := a.repo.GetArticlesBySince(ctx, g)

		if err != nil {
			return nil, xerrors.Errorf("failed to get articles from DB: %w", err)
		}

		return payload, err
	}

	if !g.Until.IsZero() && g.Since.IsZero() {
		payload, err := a.repo.GetArticlesByUntil(ctx, g)

		if err != nil {
			return nil, xerrors.Errorf("failed to get articles from DB: %w", err)
		}

		return payload, err
	}

	payload, err := a.repo.GetNewArticles(ctx, g)

	if err != nil {
		return nil, xerrors.Errorf("failed to get articles from DB: %w", err)
	}

	return payload, err
}

func (a articleModel) Create(ctx context.Context, r *http.Request, c *form.CreateArticleForm) error {
	timeStamp := time.Now().UTC().In(time.FixedZone("Asia/Tokyo", 9*60*60))

	article := entity.Article{
		ID:        uuid.NewV4().String(),
		Title: c.Title,
		URL: c.URL,
		PubDate: c.PubDate,
		CreatedAt: timeStamp,
		UpdatedAt: timeStamp,
	}

	err := a.repo.Create(r.Context(), &article)

	if err != nil {
		return xerrors.Errorf("failed to insert an article to DB: %w", err)
	}

	return err
}
