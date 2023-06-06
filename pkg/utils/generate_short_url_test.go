package utils

import "testing"

func TestUtils_GenerateShortURL(t *testing.T) {
	originalURL := "https://ozon.ru"

	short := GenerateShortURL(originalURL)

	if short != GenerateShortURL(originalURL) {
		t.Error("Generated different short URLs for the same original URL")
	}
}
