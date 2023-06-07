package handler

import (
	"errors"
	"net/http"

	"github.com/fidesy/ozon-test/internal/domain"
	"github.com/fidesy/ozon-test/internal/infrastructure/dberrors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getOriginalURL(c *gin.Context) {
	hash := c.Param("hash")
	url, err := h.service.URL.GetOriginalURL(hash)
	if err != nil {
		if errors.Is(err, dberrors.ErrHashDoesNotExist) {
			newErrorResponse(c, http.StatusNotFound, domain.ErrURLNotFound)
			return
		}

		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"original_url": url,
	})
}

func (h *Handler) createShortURL(c *gin.Context) {
	var input domain.URL

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if !h.service.URL.IsURLValid(input.OriginalURL) {
		newErrorResponse(c, http.StatusBadRequest, domain.ErrURLIsInvalid)
		return
	}

	shortURL, err := h.service.URL.CreateShortURL(input)
	if errors.Is(err, domain.ErrURLAlreadyExists) {
		newErrorResponse(c, http.StatusConflict, err)
		return
	}
	
	c.JSON(http.StatusCreated, map[string]interface{}{
		"short_url": shortURL,
	})
}
