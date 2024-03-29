package domain

import (
	"context"
)

type URL struct {
	ID          int    `json:"_" db:"id"`
	Hash        string `json:"hash" db:"hash"`
	OriginalURL string `json:"original_url" db:"original_url" binding:"required"`
}

type URLRepository interface {
	CreateURL(ctx context.Context, url URL) (int, error)
	GetURLByHash(ctx context.Context, hash string) (URL, error)
}

type URLService interface {
	CreateShortURL(url URL) (string, error)
	GetOriginalURL(shortURL string) (string, error)
	IsURLValid(url string) bool
}
