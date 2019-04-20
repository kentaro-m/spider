package article

import (
	"context"
	"github.com/satori/go.uuid"
	"golang.org/x/xerrors"
	"net/http"
	"time"
)

type ArticleModel interface {
	Get(ctx context.Context, g *GetArticleForm) ([]*Article, error)
	Create(ctx context.Context, r *http.Request, c *CreateArticleForm) error
}

func NewArticleModel(r ArticleRepository) ArticleModel {
	return &articleModel{
		repo: r,
	}
}

type articleModel struct {
	repo ArticleRepository
}

func (a articleModel) Get(ctx context.Context, g *GetArticleForm) ([]*Article, error) {
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

func (a articleModel) Create(ctx context.Context, r *http.Request, c *CreateArticleForm) error {
	timeStamp := time.Now().UTC().In(time.FixedZone("Asia/Tokyo", 9*60*60))

	article := Article{
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
