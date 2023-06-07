package postgres

import (
	"context"
	"testing"

	"github.com/fidesy/ozon-test/internal/config"
	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/fidesy/ozon-test/pkg/utils"
	"github.com/stretchr/testify/assert"
)

var (
	urls = []domain.URL{
		{OriginalURL: "https://google.com"},
		{OriginalURL: "https://amazon.com/"},
		{OriginalURL: "https://apple.com/some/path"},
	}
)

func getURLRepository(t *testing.T) *URLRepository {
	pool, err := NewPool(context.Background(), config.Default.Postgres)
	assert.Nil(t, err)

	repo := NewURLRepository(pool)

	return repo
}

func TestPostgresURLRepository_CreateURL(t *testing.T) {
	repo := getURLRepository(t)

	for i := range urls {
		urls[i].Hash = utils.GenerateShortURL(urls[i].OriginalURL)

		id, err := repo.CreateURL(
			context.Background(),
			urls[i],
		)

		assert.Nil(t, err)
		assert.NotEqual(t, 0, id)
	}
}

func TestPostgresURLRepository_GetURLByHash(t *testing.T) {
	repo := getURLRepository(t)

	for _, url := range urls {
		_url, err := repo.GetURLByHash(context.Background(), url.Hash)
		assert.Nil(t, err)
		assert.Equal(t, url.OriginalURL, _url.OriginalURL)
	}
}