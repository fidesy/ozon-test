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

	// repo = NewURLRepository()
)

// only makes sense in context with GetURL
func TestInMemoryURLRepository_CreateURL(t *testing.T) {
	repo := NewURLRepository()

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

func TestInMemoryURLRepository_GetURLByHash(t *testing.T) {
	repo := NewURLRepository()

	for i := range urls {
		// create short url
		urls[i].Hash = utils.GenerateShortURL(urls[i].OriginalURL)
		_, err := repo.CreateURL(
			context.Background(),
			urls[i],
		)
		assert.Nil(t, err)

		// get original URL from the short
		url, err := repo.GetURLByHash(context.Background(), urls[i].Hash)
		assert.Nil(t, err)
		assert.Equal(t, urls[i].OriginalURL, url.OriginalURL)
	}
}
