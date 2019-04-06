package article

import (
	"context"
	"database/sql"
	"github.com/kentaro-m/spider/api/entities"
	aRepo "github.com/kentaro-m/spider/api/repository"
	"log"
)

func NewMySQLArticleRepository(Conn *sql.DB) aRepo.ArticleRepository {
	return &mysqlArticleRepository{
		Conn: Conn,
	}
}

type mysqlArticleRepository struct {
	Conn *sql.DB
}

func (m *mysqlArticleRepository) Get(ctx context.Context) ([]*entities.Article, error) {
	query := "SELECT id, title, url, pub_date FROM articles"

	rows, err := m.Conn.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	payload := make([]*entities.Article, 0)
	for rows.Next() {
		data := new(entities.Article)
		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.URL,
			&data.PubDate,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlArticleRepository) Create(ctx context.Context, a *entities.Article) error {
	query := "INSERT INTO articles SET id = ?, title = ?, url = ?, pub_date = ?, created_at = ?, updated_at = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
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
