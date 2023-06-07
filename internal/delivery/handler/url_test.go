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
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	urls = []domain.URL{
		{OriginalURL: "https://google.com/search"},
		{OriginalURL: "https://amazon.com/products"},
		{OriginalURL: "https://apple.com/some/path/deep"},
	}
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
	router := getRouter(t)

	for i, url := range urls {
		body, _ := json.Marshal(domain.URL{
			OriginalURL: url.OriginalURL,
		})

		req, _ := http.NewRequest(http.MethodPost, "/api/create", bytes.NewBuffer(body))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		var responseBody struct {
			ShortURL string `json:"short_url"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.Nil(t, err)

		urls[i].Hash = strings.Split(responseBody.ShortURL, "/")[3]
	}
}

func TestURLHandler_getOriginalURL(t *testing.T) {

	router := getRouter(t)

	for _, url := range urls {
		req, _ := http.NewRequest(http.MethodGet, "/"+url.Hash, nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	}
}