package repository

import (
	"context"

	"github.com/kentaro-m/spider/api/entity"
)

type ArticleRepository interface {
	Get(ctx context.Context) ([]*entity.ArticleEntity, error)
	Create(ctx context.Context, a *entity.ArticleEntity) (error)
}