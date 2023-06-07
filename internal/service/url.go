package service

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/fidesy/ozon-test/internal/config"
	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/fidesy/ozon-test/pkg/utils"
)

const (
	verifyPattern = `^(https?|ftp)://[^\s/$.?#].[^\s]*$`
)

type URLServiceImpl struct {
	conf config.Config
	repo domain.URLRepository
}

func NewURLServiceImpl(conf config.Config, repo domain.URLRepository) *URLServiceImpl {
	u := &URLServiceImpl{
		conf: conf,
		repo: repo,
	}

	if u.conf.Port != "" {
		u.conf.Port = ":" + u.conf.Port
	}

	return u
}

var _ domain.URLService = &URLServiceImpl{}

func (u *URLServiceImpl) CreateShortURL(url domain.URL) (string, error) {
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
				return shortURL, domain.ErrURLAlreadyExists
			}

			// collision occurred
			sequence += hash
			continue
		}

		break
	}

	url.Hash = hash

	_, err := u.repo.CreateURL(
		context.TODO(),
		url,
	)
	if err != nil {
		log.Println("error creating short url:", err.Error())
		return "", err
	}

	return shortURL, nil
}

func (u *URLServiceImpl) GetOriginalURL(hash string) (string, error) {
	url, err := u.repo.GetURLByHash(context.TODO(), hash)
	if err != nil {
		return "", err
	}

	return url.OriginalURL, nil
}

func (u *URLServiceImpl) IsURLValid(url string) bool {
	match, err := regexp.MatchString(verifyPattern, url)
	if err != nil || !match {
		return false
	}

	return true
}
