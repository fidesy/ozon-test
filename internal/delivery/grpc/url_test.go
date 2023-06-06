package grpc

import (
	"context"
	"fmt"
	"testing"

	"github.com/fidesy/ozon-test/internal/config"
	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/fidesy/ozon-test/internal/infrastructure/persistence"
	"github.com/fidesy/ozon-test/internal/usecase"
	"github.com/fidesy/ozon-test/pkg/utils"
	shortener "github.com/fidesy/ozon-test/proto"
	"github.com/stretchr/testify/assert"
)

var (
	urls = []domain.URL{
		{OriginalURL: "https://google.com/search/some"},
		{OriginalURL: "https://amazon.com/products/phones"},
		{OriginalURL: "https://apple.com/some/path/deep/"},
	}
)

func TestGRPCURLHandler_createShortURL(t *testing.T) {
	defaultConfig := config.Default
	
	repos, err := persistence.NewRepository(context.Background(), defaultConfig)
	assert.Nil(t, err)

	usecases := usecase.NewUsecase(defaultConfig, repos)
	server := NewServer(usecases)

	for i := range urls {
		urls[i].Hash = utils.GenerateShortURL(urls[i].OriginalURL)
		expected := fmt.Sprintf("%s:%s/%s", 
			defaultConfig.Host, 
			defaultConfig.Port,
			urls[i].Hash,
		)
		request := &shortener.CreateShortURLRequest{
			OriginalUrl: urls[i].OriginalURL,
		}
		
		response, err := server.CreateShortURL(context.Background(), request)
		assert.NoError(t, err)
		assert.Equal(t, expected, response.ShortUrl)
	}
}

func TestGRPCURLHandler_getOriginalURL(t *testing.T) {
	defaultConfig := config.Default
	
	repos, err := persistence.NewRepository(context.Background(), defaultConfig)
	assert.Nil(t, err)

	usecases := usecase.NewUsecase(defaultConfig, repos)
	server := NewServer(usecases)

	for i := range urls {
		request := &shortener.GetOriginalURLRequest{
			Hash: urls[i].Hash,
		}
		
		response, err := server.GetOriginalURL(context.Background(), request)
		assert.NoError(t, err)
		assert.Equal(t, urls[i].OriginalURL, response.OriginalUrl)
	}
}
