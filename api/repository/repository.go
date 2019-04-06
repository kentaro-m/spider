package repository

import (
	"context"

	"github.com/kentaro-m/spider/api/entities"
)

type ArticleRepository interface {
	Get(ctx context.Context) ([]*entities.Article, error)
	Create(ctx context.Context, a *entities.Article) (error)
}