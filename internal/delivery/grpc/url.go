package grpc

import (
	"context"
	"errors"

	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/fidesy/ozon-test/internal/infrastructure/dberrors"
	shortener "github.com/fidesy/ozon-test/proto"
)

func (s *Server) GetOriginalURL(ctx context.Context, req *shortener.GetOriginalURLRequest) (*shortener.GetOriginalURLResponse, error) {
	hash := req.Hash

	url, err := s.service.URL.GetOriginalURL(hash)
	if err != nil {
		if errors.Is(err, dberrors.ErrHashDoesNotExist) {
			return nil, domain.ErrURLNotFound
		}
		return nil, err
	}

	response := &shortener.GetOriginalURLResponse{
		OriginalUrl: url,
	}

	return response, nil
}

func (s *Server) CreateShortURL(ctx context.Context, req *shortener.CreateShortURLRequest) (*shortener.CreateShortURLResponse, error) {
	var input domain.URL

	input.OriginalURL = req.OriginalUrl

	if !s.service.URL.IsURLValid(input.OriginalURL) {
		return nil, domain.ErrURLIsInvalid
	}

	shortURL, err := s.service.URL.CreateShortURL(input)
	if errors.Is(err, domain.ErrURLAlreadyExists) {
		return nil, domain.ErrURLAlreadyExists
	}

	response := &shortener.CreateShortURLResponse{
		ShortUrl: shortURL,
	}

	return response, nil
}
