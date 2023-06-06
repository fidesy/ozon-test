package usecase

import (
	"github.com/fidesy/ozon-test/internal/config"
	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/fidesy/ozon-test/internal/infrastructure/persistence"
)

type Usecase struct {
	URL domain.URLUsecase
}

func NewUsecase(conf config.Config, repos persistence.Repository) *Usecase {
	return &Usecase{
		NewURLUsecaseImpl(conf, repos),
	}
}
