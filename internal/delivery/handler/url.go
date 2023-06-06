package handler

import (
	"net/http"

	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getOriginalURL(c *gin.Context) {
	hash := c.Param("hash")
	url, err := h.usecases.URL.GetOriginalURL(hash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"original_url": url,
	})
}

func (h *Handler) createShortURL(c *gin.Context) {
	var input domain.URL

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	shortURL := h.usecases.URL.CreateShortURL(input)

	c.JSON(http.StatusCreated, map[string]interface{}{
		"short_url": shortURL,
	})
}
