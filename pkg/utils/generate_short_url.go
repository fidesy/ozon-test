package utils

import (
	"crypto/md5"
	"encoding/base64"
	"strings"
)

const (
	LENGTH = 10
	SIGNS  = "/+-"
)

func GenerateShortURL(originalURL string) string {
	hash := md5.Sum([]byte(originalURL))

	encoded := base64.URLEncoding.EncodeToString(hash[:])

	for i := range SIGNS {
		encoded = strings.ReplaceAll(encoded, string(SIGNS[i]), "_")
	}
	encoded = strings.TrimRight(encoded, "=")

	if len(encoded) > LENGTH {
		encoded = encoded[:LENGTH]
	}

	return encoded
}
