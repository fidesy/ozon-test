package service

import (
	"github.com/fidesy/ozon-test/internal/config"
	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/fidesy/ozon-test/internal/infrastructure/persistence"
)

type Service struct {
	URL domain.URLService
}

func NewService(conf config.Config, repos persistence.Repository) *Service {
	return &Service{
		NewURLServiceImpl(conf, repos),
	}
}
