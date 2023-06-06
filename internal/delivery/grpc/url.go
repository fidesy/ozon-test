package grpc

import (
	"context"

	"github.com/fidesy/ozon-test/internal/domain"
	shortener "github.com/fidesy/ozon-test/proto"
)

func (s *Server) GetOriginalURL(ctx context.Context, req *shortener.GetOriginalURLRequest) (*shortener.GetOriginalURLResponse, error) {
	urlStr, err := s.usecases.URL.GetOriginalURL(req.Hash)
	if err != nil {
		return nil, err
	}

	response := &shortener.GetOriginalURLResponse{
		OriginalUrl: urlStr,
	}

	return response, nil
}

func (s *Server) CreateShortURL(ctx context.Context, req *shortener.CreateShortURLRequest) (*shortener.CreateShortURLResponse, error) {
	input := domain.URL{
		OriginalURL: req.OriginalUrl,
	}
	shortURL := s.usecases.URL.CreateShortURL(input)

	response := &shortener.CreateShortURLResponse{
		ShortUrl: shortURL,
	}

	return response, nil
}
