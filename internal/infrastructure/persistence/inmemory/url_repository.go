package inmemory

import (
	"context"
	"sync"

	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/fidesy/ozon-test/internal/infrastructure/errors"
)

type URLRepository struct {
	hashToURL map[string]domain.URL
	sync.RWMutex
}

func NewURLRepository() *URLRepository {
	return &URLRepository{
		hashToURL: make(map[string]domain.URL),
	}
}

var _ domain.URLRepository = &URLRepository{}

func (r *URLRepository) CreateURL(ctx context.Context, url domain.URL) (int, error) {
	r.RLock()
	defer r.RUnlock()

	url.ID = len(r.hashToURL) + 1
	r.hashToURL[url.Hash] = url

	return url.ID, nil
}

func (r *URLRepository) GetURLByHash(ctx context.Context, hash string) (domain.URL, error) {
	r.RLock()
	defer r.RUnlock()

	url, ok := r.hashToURL[hash]
	if !ok {
		return domain.URL{}, errors.ErrHashDoesNotExist
	}

	return url, nil
}
