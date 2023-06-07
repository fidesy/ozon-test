package inmemory

import "github.com/fidesy/ozon-test/internal/domain"

type InMemory struct {
	domain.URLRepository
}

func New() *InMemory {
	return &InMemory{
		NewURLRepository(),
	}
}

// dummy function to implement the interface
func (i *InMemory) Close() error {
	return nil
}