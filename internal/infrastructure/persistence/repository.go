package persistence

import (
	"context"

	"github.com/fidesy/ozon-test/internal/config"
	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/fidesy/ozon-test/internal/infrastructure/errors"
	"github.com/fidesy/ozon-test/internal/infrastructure/persistence/inmemory"
	"github.com/fidesy/ozon-test/internal/infrastructure/persistence/postgres"
)

type Repository interface {
	domain.URLRepository
}

func NewRepository(ctx context.Context, conf config.Config) (Repository, error) {
	switch conf.Database {
	case "postgres":
		pool, err := postgres.NewPool(ctx, conf.Postgres)
		if err != nil {
			return nil, err
		}

		return postgres.New(pool), nil
	case "in-memory":
		return inmemory.New(), nil
	default:
		return nil, errors.ErrInvalidDatabaseName
	}
}
