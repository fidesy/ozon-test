package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fidesy/ozon-test/internal/config"
	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/fidesy/ozon-test/internal/infrastructure/persistence"
	"github.com/fidesy/ozon-test/internal/service"
	"github.com/fidesy/ozon-test/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func getRouter(t *testing.T) *gin.Engine {
	conf := config.Default

	repos, err := persistence.NewRepository(context.Background(), config.Default)
	assert.Nil(t, err)

	service := service.NewService(conf, repos)
	handler := New(service)

	return handler.InitRoutes()
}

func TestURLHandler_createShortURL(t *testing.T) {
	tests := []struct {
		URL        domain.URL
		StatusCode int
	}{
		{
			URL:        domain.URL{OriginalURL: "https://google.com/search"},
			StatusCode: http.StatusCreated,
		},
		{
			URL:        domain.URL{OriginalURL: "https://google.com/search"},
			StatusCode: http.StatusConflict,
		},
		{
			URL:        domain.URL{OriginalURL: "http//google.com"},
			StatusCode: http.StatusBadRequest,
		},
		{
			URL:        domain.URL{OriginalURL: "https://google.com/search/some/path"},
			StatusCode: http.StatusCreated,
		},
	}

	router := getRouter(t)

	for _, test := range tests {
		body, _ := json.Marshal(test.URL)

		req, _ := http.NewRequest(http.MethodPost, "/api/create", bytes.NewBuffer(body))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, test.StatusCode, w.Code)

		var responseBody struct {
			ShortURL string `json:"short_url"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.Nil(t, err)

		urlSlice := strings.Split(responseBody.ShortURL, "/")
		if len(urlSlice) == 1 {
			continue
		}

		hash := urlSlice[3]
		assert.Equal(t, utils.GenerateShortURL(test.URL.OriginalURL), hash)
	}
}

func TestURLHandler_getOriginalURL(t *testing.T) {
	tests := []struct {
		URL        domain.URL
		StatusCode int
	}{
		{
			URL:        domain.URL{OriginalURL: "https://google.com/search/query"},
			StatusCode: http.StatusOK,
		},
		{
			URL:        domain.URL{OriginalURL: "http://website.org"},
			StatusCode: http.StatusOK,
		},
	}

	router := getRouter(t)

	// create short URLs
	for i, test := range tests {
		body, _ := json.Marshal(test.URL)

		req, _ := http.NewRequest(http.MethodPost, "/api/create", bytes.NewBuffer(body))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		var responseBody struct {
			ShortURL string `json:"short_url"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.Nil(t, err)

		hash := strings.Split(responseBody.ShortURL, "/")[3]
		assert.Equal(t, utils.GenerateShortURL(test.URL.OriginalURL), hash)

		tests[i].URL.Hash = hash
	}

	// get original URL from the short
	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodGet, "/"+test.URL.Hash, nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, test.StatusCode, w.Code)
	}

	// test not found
	req, _ := http.NewRequest(http.MethodGet, "/neverexisted", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
