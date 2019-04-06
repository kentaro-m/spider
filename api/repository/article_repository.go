package repository

import (
	"context"
	"database/sql"
	"github.com/kentaro-m/spider/api/entity"
	"log"
)

func NewArticleRepository(Conn *sql.DB) ArticleRepository {
	return &articleRepository{
		Conn: Conn,
	}
}

type ArticleRepository interface {
	Get(ctx context.Context) ([]*entity.Article, error)
	Create(ctx context.Context, a *entity.Article) (error)
}

type articleRepository struct {
	Conn *sql.DB
}

func (ar articleRepository) Get(ctx context.Context) ([]*entity.Article, error) {
	query := "SELECT id, title, url, pub_date, created_at, updated_at FROM articles"

	rows, err := ar.Conn.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

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
			log.Fatal(err)
			return nil, err
		}

		payload = append(payload, data)
	}
	return payload, nil
}

func (ar *articleRepository) Create(ctx context.Context, a *entity.Article) error {
	query := "INSERT INTO articles SET id = ?, title = ?, url = ?, pub_date = ?, created_at = ?, updated_at = ?"

	stmt, err := ar.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, a.ID, a.Title, a.URL, a.PubDate, a.CreatedAt, a.UpdatedAt)
	defer stmt.Close()

	if err != nil {
		return err
	}

	return nil
}
