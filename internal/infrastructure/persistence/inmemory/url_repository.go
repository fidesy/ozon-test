package inmemory

import (
	"context"
	"sync"

	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/fidesy/ozon-test/internal/infrastructure/dberrors"
)

type URLRepository struct {
	hashToURL map[string]domain.URL
	counter   int
	sync.RWMutex
}

func NewURLRepository() *URLRepository {
	return &URLRepository{
		hashToURL: make(map[string]domain.URL),
	}
}

var _ domain.URLRepository = &URLRepository{}

func (r *URLRepository) CreateURL(ctx context.Context, url domain.URL) (int, error) {
	r.Lock()
	defer r.Unlock()

	r.counter++
	url.ID = r.counter
	r.hashToURL[url.Hash] = url

	return url.ID, nil
}

func (r *URLRepository) GetURLByHash(ctx context.Context, hash string) (domain.URL, error) {
	r.RLock()
	defer r.RUnlock()

	url, ok := r.hashToURL[hash]
	if !ok {
		return domain.URL{}, dberrors.ErrHashDoesNotExist
	}

	return url, nil
}
