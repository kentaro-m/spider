package article

import (
	"context"
	s "github.com/kentaro-m/spider/api/site"
	"github.com/satori/go.uuid"
	"golang.org/x/xerrors"
	"net/http"
	"time"
)

type ArticleModel interface {
	Get(ctx context.Context, g *GetArticleForm) ([]*Item, error)
	Create(ctx context.Context, r *http.Request, c *CreateArticleForm) error
}

func NewArticleModel(ar ArticleRepository, sr s.SiteRepository) ArticleModel {
	return &articleModel{
		aRepo: ar,
		sRepo: sr,
	}
}

type articleModel struct {
	aRepo ArticleRepository
	sRepo s.SiteRepository
}

func (a articleModel) Get(ctx context.Context, g *GetArticleForm) ([]*Item, error) {
	if !g.Until().IsZero() && !g.Since().IsZero() {
		payload, err := a.aRepo.GetArticlesBySinceAndUntil(ctx, g)

		if err != nil {
			return nil, xerrors.Errorf("failed to get articles from DB: %w", err)
		}

		return payload, err
	}

	if g.Until().IsZero() && !g.Since().IsZero() {
		payload, err := a.aRepo.GetArticlesBySince(ctx, g)

		if err != nil {
			return nil, xerrors.Errorf("failed to get articles from DB: %w", err)
		}

		return payload, err
	}

	if !g.Until().IsZero() && g.Since().IsZero() {
		payload, err := a.aRepo.GetArticlesByUntil(ctx, g)

		if err != nil {
			return nil, xerrors.Errorf("failed to get articles from DB: %w", err)
		}

		return payload, err
	}

	payload, err := a.aRepo.GetNewArticles(ctx, g)

	if err != nil {
		return nil, xerrors.Errorf("failed to get articles from DB: %w", err)
	}

	return payload, err
}

func (a articleModel) Create(ctx context.Context, r *http.Request, c *CreateArticleForm) error {
	timeStamp := time.Now().UTC().In(time.FixedZone("Asia/Tokyo", 9*60*60))

	getSiteForm := s.GetSiteForm{
		Title: c.SiteTitle(),
	}

	data, err := a.sRepo.GetSiteInfo(r.Context(), &getSiteForm)

	if err != nil {
		return xerrors.Errorf("failed to get site info from DB: %w", err)
	}

	siteID := data.ID

	if siteID == "" {
		site := s.Site{
			ID: uuid.NewV4().String(),
			Title: c.SiteTitle(),
			URL: c.SiteURL(),
			CreatedAt: timeStamp,
			UpdatedAt: timeStamp,
		}

		siteID, err = a.sRepo.Create(r.Context(), &site)

		if err != nil {
			return xerrors.Errorf("failed to insert site info to DB: %w", err)
		}
	}

	article := Article{
		article{
			ID:        uuid.NewV4().String(),
			Title:     c.Title(),
			URL:       c.URL(),
			PubDate:   c.PubDate(),
			SiteID:    siteID,
			CreatedAt: timeStamp,
			UpdatedAt: timeStamp,
		},
	}

	err = a.aRepo.Create(r.Context(), &article)

	if err != nil {
		return xerrors.Errorf("failed to insert an article to DB: %w", err)
	}

	return err
}
