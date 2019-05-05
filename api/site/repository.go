package site

import (
	"context"
	"database/sql"
	"golang.org/x/xerrors"
)

func NewSiteRepository(Conn *sql.DB) SiteRepository {
	return &siteRepository{
		Conn: Conn,
	}
}

type SiteRepository interface {
	GetSiteInfo(ctx context.Context, g *GetSiteForm) (*Site, error)
	Create(ctx context.Context, s *Site) (string, error)
}

type siteRepository struct {
	Conn *sql.DB
}

func (sr *siteRepository) GetSiteInfo(ctx context.Context, g *GetSiteForm) (*Site, error) {
	query := "SELECT id, title, url, created_at, updated_at FROM sites WHERE title = ?"

	rows, err := sr.Conn.QueryContext(ctx, query, g.Title)

	if err != nil {
		return nil, xerrors.Errorf("failed to execute a query: %w", err)
	}

	defer func() {
		er := rows.Close()

		if er != nil {
			err = xerrors.Errorf("failed to close a connection: %w", er)
		}
	}()

	payload := new(Site)
	for rows.Next() {
		err := rows.Scan(
			&payload.ID,
			&payload.Title,
			&payload.URL,
			&payload.CreatedAt,
			&payload.UpdatedAt,
		)

		if err != nil {
			return nil, xerrors.Errorf("failed to convert columns into Go types: %w", err)
		}
	}

	return payload, err
}

func (sr *siteRepository) Create(ctx context.Context, s *Site) (string, error) {
	query := "INSERT INTO sites SET ID = ?, title = ?, url = ?, created_at = ?, updated_at = ?"

	stmt, err := sr.Conn.PrepareContext(ctx, query)

	if err != nil {
		return "", xerrors.Errorf("failed to create a prepared statement: %w", err)
	}

	_, err = stmt.ExecContext(ctx, s.ID, s.Title, s.URL, s.CreatedAt.Format("2006-01-02 15:04:05"), s.UpdatedAt.Format("2006-01-02 15:04:05"))

	if err != nil {
		return "", xerrors.Errorf("failed to execute a prepared statement: %w", err)
	}

	defer func() {
		er := stmt.Close()

		if er != nil {
			err = xerrors.Errorf("failed to close a connection: %w", er)
		}
	}()

	return s.ID, err
}