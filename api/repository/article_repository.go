package repository

import (
	"context"
	"database/sql"
	"github.com/kentaro-m/spider/api/entity"
)

func NewArticleRepository(Conn *sql.DB) ArticleRepository {
	return &articleRepository{
		Conn: Conn,
	}
}

type ArticleRepository interface {
	Get(ctx context.Context) ([]*entity.Article, error)
	Create(ctx context.Context, a *entity.Article) error
}

type articleRepository struct {
	Conn *sql.DB
}

func (ar articleRepository) Get(ctx context.Context) ([]*entity.Article, error) {
	query := "SELECT id, title, url, pub_date, created_at, updated_at FROM articles"

	rows, err := ar.Conn.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer func() {
		er := rows.Close()

		if er != nil {
			err = er
		}
	}()

	payload := make([]*entity.Article, 0)
	for rows.Next() {
		data := new(entity.Article)
		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.URL,
			&data.PubDate,
			&data.CreatedAt,
			&data.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		payload = append(payload, data)
	}

	return payload, err
}

func (ar *articleRepository) Create(ctx context.Context, a *entity.Article) error {
	query := "INSERT INTO articles SET id = ?, title = ?, url = ?, pub_date = ?, created_at = ?, updated_at = ?"

	stmt, err := ar.Conn.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, a.ID, a.Title, a.URL, a.PubDate.Format("2006-01-02 15:04:05"), a.CreatedAt.Format("2006-01-02 15:04:05"), a.UpdatedAt.Format("2006-01-02 15:04:05"))

	if err != nil {
		return err
	}

	defer func() {
		er := stmt.Close()

		if er != nil {
			err = er
		}
	}()

	return err
}
