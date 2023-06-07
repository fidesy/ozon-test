package postgres

import (
	"context"
	"fmt"

	"github.com/fidesy/ozon-test/internal/config"
	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool
	domain.URLRepository
}

func New(pool *pgxpool.Pool) *Postgres {
	return &Postgres{
		pool: pool,
		URLRepository: NewURLRepository(pool),
	}
}

func NewPool(ctx context.Context, conf config.Postgres) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(
		ctx,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.DBName,
			conf.SSLMode,
		))
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}

func (p *Postgres) Close() error {
	p.pool.Close()
	return nil
}