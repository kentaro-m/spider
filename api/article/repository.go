package article

import (
	"context"
	"database/sql"
	"fmt"
	"golang.org/x/xerrors"
	"time"
)

func NewArticleRepository(Conn *sql.DB) ArticleRepository {
	return &articleRepository{
		Conn: Conn,
	}
}

type ArticleRepository interface {
	GetNewArticles(ctx context.Context, g *GetArticleForm) ([]*Item, error)
	GetArticlesBySince(ctx context.Context, g *GetArticleForm) ([]*Item, error)
	GetArticlesByUntil(ctx context.Context, g *GetArticleForm) ([]*Item, error)
	GetArticlesBySinceAndUntil(ctx context.Context, g *GetArticleForm) ([]*Item, error)
	Create(ctx context.Context, a *Article) error
}

type articleRepository struct {
	Conn *sql.DB
}

type Item struct {
	ID        	string    `json:"id" example:"faf9c3a7-b3ee-441f-baec-a5b668948381"`
	Title     	string    `json:"title" example:"AWS CDKでサーバーレスアプリケーションのデプロイを試す"`
	URL       	string    `json:"url" example:"https://blog.kentarom.com/learn-aws-cdk/"`
	PubDate   	time.Time `json:"pub_date" example:"2019-01-19T14:13:01Z"`
	SiteTitle   string    `json:"site_title" example:"AWS CDKでサーバーレスアプリケーションのデプロイを試す"`
	SiteURL     string    `json:"site_url" example:"https://blog.kentarom.com/learn-aws-cdk/"`
}

func (ar articleRepository) GetNewArticles(ctx context.Context, g *GetArticleForm) ([]*Item, error) {
	query := fmt.Sprintf(`
		SELECT a.id, a.title, a.url, a.pub_date, s.title AS site_title, s.url AS site_url
		FROM articles AS a
		LEFT OUTER JOIN sites AS s ON a.site_id = s.id
		ORDER BY pub_date DESC
		LIMIT ?
	`)

	if g.Sort() == "asc" {
		query = fmt.Sprintf(`
			SELECT a.id, a.title, a.url, a.pub_date, s.title AS site_title, s.url AS site_url
			FROM articles AS a
			LEFT OUTER JOIN sites AS s ON a.site_id = s.id
			ORDER BY pub_date ASC
			LIMIT ?
		`)
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

	payload := make([]*Item, 0)
	for rows.Next() {
		data := new(Item)
		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.URL,
			&data.PubDate,
			&data.SiteTitle,
			&data.SiteURL,
		)

		if err != nil {
			return nil, xerrors.Errorf("failed to convert columns into Go types: %w", err)
		}

		payload = append(payload, data)
	}

	return payload, err
}

func (ar articleRepository) GetArticlesBySince(ctx context.Context, g *GetArticleForm) ([]*Item, error) {
	query := fmt.Sprintf(`
		SELECT a.id, a.title, a.url, a.pub_date, s.title AS site_title, s.url AS site_url
		FROM articles AS a
		LEFT OUTER JOIN sites AS s ON a.site_id = s.id
		WHERE pub_date >= ?
		ORDER BY pub_date DESC
		LIMIT ?
	`)

	if g.Sort() == "asc" {
		query = fmt.Sprintf(`
			SELECT a.id, a.title, a.url, a.pub_date, s.title AS site_title, s.url AS site_url
			FROM articles AS a
			LEFT OUTER JOIN sites AS s ON a.site_id = s.id
			WHERE pub_date >= ?
			ORDER BY pub_date ASC
			LIMIT ?
		`)
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

	payload := make([]*Item, 0)
	for rows.Next() {
		data := new(Item)
		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.URL,
			&data.PubDate,
			&data.SiteTitle,
			&data.SiteURL,
		)

		if err != nil {
			return nil, xerrors.Errorf("failed to convert columns into Go types: %w", err)
		}

		payload = append(payload, data)
	}

	return payload, err
}

func (ar articleRepository) GetArticlesByUntil(ctx context.Context, g *GetArticleForm) ([]*Item, error) {
	query := fmt.Sprintf(`
		SELECT a.id, a.title, a.url, a.pub_date, s.title AS site_title, s.url AS site_url
		FROM articles AS a
		LEFT OUTER JOIN sites AS s ON a.site_id = s.id
		WHERE pub_date <= ?
		ORDER BY pub_date DESC
		LIMIT ?
	`)

	if g.Sort() == "asc" {
		query = fmt.Sprintf(`
			SELECT a.id, a.title, a.url, a.pub_date, s.title AS site_title, s.url AS site_url
			FROM articles AS a
			LEFT OUTER JOIN sites AS s ON a.site_id = s.id
			WHERE pub_date <= ?
			ORDER BY pub_date ASC
			LIMIT ?
		`)
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

	payload := make([]*Item, 0)
	for rows.Next() {
		data := new(Item)
		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.URL,
			&data.PubDate,
			&data.SiteTitle,
			&data.SiteURL,
		)

		if err != nil {
			return nil, xerrors.Errorf("failed to convert columns into Go types: %w", err)
		}

		payload = append(payload, data)
	}

	return payload, err
}

func (ar articleRepository) GetArticlesBySinceAndUntil(ctx context.Context, g *GetArticleForm) ([]*Item, error) {
	query := fmt.Sprintf(`
		SELECT a.id, a.title, a.url, a.pub_date, s.title AS site_title, s.url AS site_url
		FROM articles AS a
		LEFT OUTER JOIN sites AS s ON a.site_id = s.id
		WHERE pub_date >= ? AND pub_date <= ?
		ORDER BY pub_date DESC
		LIMIT ?
	`)

	if g.Sort() == "asc" {
		query = fmt.Sprintf(`
			SELECT a.id, a.title, a.url, a.pub_date, s.title AS site_title, s.url AS site_url
			FROM articles AS a
			LEFT OUTER JOIN sites AS s ON a.site_id = s.id
			WHERE pub_date >= ? AND pub_date <= ?
			ORDER BY pub_date ASC
			LIMIT ?
		`)
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

	payload := make([]*Item, 0)
	for rows.Next() {
		data := new(Item)
		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.URL,
			&data.PubDate,
			&data.SiteTitle,
			&data.SiteURL,
		)

		if err != nil {
			return nil, xerrors.Errorf("failed to convert columns into Go types: %w", err)
		}

		payload = append(payload, data)
	}

	return payload, err
}

func (ar *articleRepository) Create(ctx context.Context, a *Article) error {
	query := fmt.Sprintf(`
		INSERT INTO articles
		SET ID = ?, title = ?, url = ?, pub_date = ?, site_id = ?, created_at = ?, updated_at = ?
	`)

	stmt, err := ar.Conn.PrepareContext(ctx, query)

	if err != nil {
		return xerrors.Errorf("failed to create a prepared statement: %w", err)
	}

	_, err = stmt.ExecContext(ctx, a.ID(), a.Title(), a.URL(), a.PubDate().Format("2006-01-02 15:04:05"), a.SiteID(), a.CreatedAt().Format("2006-01-02 15:04:05"), a.UpdatedAt().Format("2006-01-02 15:04:05"))

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
