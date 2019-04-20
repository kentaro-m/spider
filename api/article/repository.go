package article

import (
	"context"
	"database/sql"
	"golang.org/x/xerrors"
)

func NewArticleRepository(Conn *sql.DB) ArticleRepository {
	return &articleRepository{
		Conn: Conn,
	}
}

type ArticleRepository interface {
	GetNewArticles(ctx context.Context, g *GetArticleForm) ([]*Article, error)
	GetArticlesBySince(ctx context.Context, g *GetArticleForm) ([]*Article, error)
	GetArticlesByUntil(ctx context.Context, g *GetArticleForm) ([]*Article, error)
	GetArticlesBySinceAndUntil(ctx context.Context, g *GetArticleForm) ([]*Article, error)
	Create(ctx context.Context, a *Article) error
}

type articleRepository struct {
	Conn *sql.DB
}

func (ar articleRepository) GetNewArticles(ctx context.Context, g *GetArticleForm) ([]*Article, error) {
	query := "SELECT ID, title, url, pub_date, created_at, updated_at FROM articles ORDER BY pub_date DESC LIMIT ?"

	if g.Sort() == "asc" {
		query = "SELECT ID, title, url, pub_date, created_at, updated_at FROM articles ORDER BY pub_date ASC LIMIT ?"
	}

	rows, err := ar.Conn.QueryContext(ctx, query, g.Limit())

	if err != nil {
		return nil, xerrors.Errorf("failed to execute a query: %w", err)
	}

	defer func() {
		er := rows.Close()

		if er != nil {
			err = xerrors.Errorf("failed to close a connection: %w", er)
		}
	}()

	payload := make([]*Article, 0)
	for rows.Next() {
		data := new(Article)
		err := rows.Scan(
			&data.article.ID,
			&data.article.Title,
			&data.article.URL,
			&data.article.PubDate,
			&data.article.CreatedAt,
			&data.article.UpdatedAt,
		)

		if err != nil {
			return nil, xerrors.Errorf("failed to convert columns into Go types: %w", err)
		}

		payload = append(payload, data)
	}

	return payload, err
}

func (ar articleRepository) GetArticlesBySince(ctx context.Context, g *GetArticleForm) ([]*Article, error) {
	query := "SELECT ID, title, url, pub_date, created_at, updated_at FROM articles WHERE pub_date >= ? ORDER BY pub_date DESC LIMIT ?"

	if g.Sort() == "asc" {
		query = "SELECT ID, title, url, pub_date, created_at, updated_at FROM articles WHERE pub_date >= ? ORDER BY pub_date ASC LIMIT ?"
	}

	rows, err := ar.Conn.QueryContext(ctx, query, g.Since().Format("2006-01-02 15:04:05"), g.Limit())

	if err != nil {
		return nil, xerrors.Errorf("failed to execute a query: %w", err)
	}

	defer func() {
		er := rows.Close()

		if er != nil {
			err = xerrors.Errorf("failed to close a connection: %w", er)
		}
	}()

	payload := make([]*Article, 0)
	for rows.Next() {
		data := new(Article)
		err := rows.Scan(
			&data.article.ID,
			&data.article.Title,
			&data.article.URL,
			&data.article.PubDate,
			&data.article.CreatedAt,
			&data.article.UpdatedAt,
		)

		if err != nil {
			return nil, xerrors.Errorf("failed to convert columns into Go types: %w", err)
		}

		payload = append(payload, data)
	}

	return payload, err
}

func (ar articleRepository) GetArticlesByUntil(ctx context.Context, g *GetArticleForm) ([]*Article, error) {
	query := "SELECT ID, title, url, pub_date, created_at, updated_at FROM articles WHERE pub_date <= ? ORDER BY pub_date DESC LIMIT ?"

	if g.Sort() == "asc" {
		query = "SELECT ID, title, url, pub_date, created_at, updated_at FROM articles WHERE pub_date <= ? ORDER BY pub_date ASC LIMIT ?"
	}

	rows, err := ar.Conn.QueryContext(ctx, query, g.Until().Format("2006-01-02 15:04:05"), g.Limit())

	if err != nil {
		return nil, xerrors.Errorf("failed to execute a query: %w", err)
	}

	defer func() {
		er := rows.Close()

		if er != nil {
			err = xerrors.Errorf("failed to close a connection: %w", er)
		}
	}()

	payload := make([]*Article, 0)
	for rows.Next() {
		data := new(Article)
		err := rows.Scan(
			&data.article.ID,
			&data.article.Title,
			&data.article.URL,
			&data.article.PubDate,
			&data.article.CreatedAt,
			&data.article.UpdatedAt,
		)

		if err != nil {
			return nil, xerrors.Errorf("failed to convert columns into Go types: %w", err)
		}

		payload = append(payload, data)
	}

	return payload, err
}

func (ar articleRepository) GetArticlesBySinceAndUntil(ctx context.Context, g *GetArticleForm) ([]*Article, error) {
	query := "SELECT ID, title, url, pub_date, created_at, updated_at FROM articles WHERE pub_date >= ? AND pub_date <= ? ORDER BY pub_date DESC LIMIT ?"

	if g.Sort() == "asc" {
		query = "SELECT ID, title, url, pub_date, created_at, updated_at FROM articles WHERE pub_date >= ? AND pub_date <= ? ORDER BY pub_date ASC LIMIT ?"
	}

	rows, err := ar.Conn.QueryContext(ctx, query, g.Since().Format("2006-01-02 15:04:05"), g.Until().Format("2006-01-02 15:04:05"), g.Limit())

	if err != nil {
		return nil, xerrors.Errorf("failed to execute a query: %w", err)
	}

	defer func() {
		er := rows.Close()

		if er != nil {
			err = xerrors.Errorf("failed to close a connection: %w", er)
		}
	}()

	payload := make([]*Article, 0)
	for rows.Next() {
		data := new(Article)
		err := rows.Scan(
			&data.article.ID,
			&data.article.Title,
			&data.article.URL,
			&data.article.PubDate,
			&data.article.CreatedAt,
			&data.article.UpdatedAt,
		)

		if err != nil {
			return nil, xerrors.Errorf("failed to convert columns into Go types: %w", err)
		}

		payload = append(payload, data)
	}

	return payload, err
}

func (ar *articleRepository) Create(ctx context.Context, a *Article) error {
	query := "INSERT INTO articles SET ID = ?, title = ?, url = ?, pub_date = ?, created_at = ?, updated_at = ?"

	stmt, err := ar.Conn.PrepareContext(ctx, query)

	if err != nil {
		return xerrors.Errorf("failed to create a prepared statement: %w", err)
	}

	_, err = stmt.ExecContext(ctx, a.ID(), a.Title(), a.URL(), a.PubDate().Format("2006-01-02 15:04:05"), a.CreatedAt().Format("2006-01-02 15:04:05"), a.UpdatedAt().Format("2006-01-02 15:04:05"))

	if err != nil {
		return xerrors.Errorf("failed to execute a prepared statement: %w", err)
	}

	defer func() {
		er := stmt.Close()

		if er != nil {
			err = xerrors.Errorf("failed to close a connection: %w", er)
		}
	}()

	return err
}
