package inmemory

import (
	"context"
	"testing"

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

	repo = NewURLRepository()
)

func TestInMemoryURLRepository_CreateURL(t *testing.T) {
	for i := range urls {
		urls[i].Hash = utils.GenerateShortURL(urls[i].OriginalURL)

		id, err := repo.CreateURL(
			context.Background(),
			domain.URL{
				OriginalURL: urls[i].OriginalURL,
				Hash:        urls[i].Hash,
			})

		assert.Nil(t, err)
		assert.NotEqual(t, 0, id)
	}
}

func TestInMemoryURLRepository_GetURLByHash(t *testing.T) {
	for _, url := range urls {
		_url, err := repo.GetURLByHash(context.Background(), url.Hash)
		assert.Nil(t, err)
		assert.Equal(t, url.OriginalURL, _url.OriginalURL)
	}
}
