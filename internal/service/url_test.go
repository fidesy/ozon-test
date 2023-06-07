package service

import (
	"context"
	"testing"

	"github.com/fidesy/ozon-test/internal/config"
	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/fidesy/ozon-test/internal/infrastructure/persistence"
	"github.com/fidesy/ozon-test/pkg/utils"
	"github.com/stretchr/testify/assert"
)


var (
	urls = []domain.URL{
		{OriginalURL: "https://yandex.ru"},
		{OriginalURL: "https://mazon.com/"},
		{OriginalURL: "https://books.com"},
	}
)

func TestURLService_CreateShortURL(t *testing.T) {
	repos, err := persistence.NewRepository(context.Background(), config.Default)
	assert.Nil(t, err)

	service := NewURLServiceImpl(config.Default, repos)

	for i := range urls {
		hash := utils.GenerateShortURL(urls[i].OriginalURL)

		shortURL, err := service.CreateShortURL(urls[i])
		assert.Nil(t, err)
		assert.Contains(t, shortURL, hash)
	}
}


func TestURLService_GetOriginalURL(t *testing.T) {
	repos, err := persistence.NewRepository(context.Background(), config.Default)
	assert.Nil(t, err)

	service := NewURLServiceImpl(config.Default, repos)

	// create URLs
	for i := range urls {
		urls[i].OriginalURL += "/additional/path"
		hash := utils.GenerateShortURL(urls[i].OriginalURL)
		urls[i].Hash = hash
		shortURL, err := service.CreateShortURL(urls[i])
		assert.Nil(t, err)
		assert.Contains(t, shortURL, hash)
	}

	for i := range urls {
		originalURL, err := service.GetOriginalURL(urls[i].Hash)
		assert.Nil(t, err)
		assert.Equal(t, urls[i].OriginalURL, originalURL)
	}
}

