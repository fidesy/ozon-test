package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/fidesy/ozon-test/internal/config"
	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/fidesy/ozon-test/pkg/utils"
)

type URLUsecaseImpl struct {
	conf config.Config
	repo domain.URLRepository
}

func NewURLUsecaseImpl(conf config.Config, repo domain.URLRepository) *URLUsecaseImpl {
	u := &URLUsecaseImpl{
		conf: conf,
		repo: repo,
	}

	if u.conf.Port != "" {
		u.conf.Port = ":" + u.conf.Port
	}

	return u
}

var _ domain.URLUsecase = &URLUsecaseImpl{}

func (u *URLUsecaseImpl) CreateShortURL(url domain.URL) string {
	var (
		hash, shortURL string
		sequence       = url.OriginalURL
	)

	for {
		hash = utils.GenerateShortURL(sequence)
		shortURL = fmt.Sprintf("%s%s/%s", u.conf.Host, u.conf.Port, hash)

		dbURL, err := u.repo.GetURLByHash(context.TODO(), hash)
		// hash already exists
		if err == nil {
			if dbURL.OriginalURL == url.OriginalURL {
				return shortURL
			}

			// collision occured
			sequence += hash
			continue
		}

		break
	}

	log.Println("HERE", shortURL)

	url.Hash = hash
	_, err := u.repo.CreateURL(
		context.TODO(),
		url,
	)
	if err != nil {
		log.Println("error creating short url:", err.Error())
		return ""
	}

	return shortURL
}

func (u *URLUsecaseImpl) GetOriginalURL(hash string) (string, error) {
	url, err := u.repo.GetURLByHash(context.TODO(), hash)
	if err != nil {
		return "", err
	}

	return url.OriginalURL, nil
}
