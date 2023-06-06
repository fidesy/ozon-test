package postgres

import (
	"context"

	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type URLRepository struct {
	pool *pgxpool.Pool
}

func NewURLRepository(pool *pgxpool.Pool) *URLRepository {
	return &URLRepository{pool}

}

var _ domain.URLRepository = &URLRepository{}

func (r *URLRepository) CreateURL(ctx context.Context, url domain.URL) (int, error) {
	var id int
	err := r.pool.QueryRow(
		ctx,
		"INSERT INTO urls(hash, original_url) VALUES($1, $2) RETURNING id",
		url.Hash,
		url.OriginalURL,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *URLRepository) GetURLByHash(ctx context.Context, hash string) (domain.URL, error) {
	var url domain.URL
	err := r.pool.QueryRow(
		ctx,
		"SELECT id, hash, original_url FROM urls WHERE hash=$1",
		hash,
	).Scan(
		&url.ID,
		&url.Hash,
		&url.OriginalURL,
	)
	if err != nil {
		return domain.URL{}, err
	}

	return url, nil
}
