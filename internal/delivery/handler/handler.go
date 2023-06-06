package handler

import (
	"github.com/fidesy/ozon-test/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecases *usecase.Usecase
}

func New(usecases *usecase.Usecase) *Handler {
	return &Handler{usecases}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.GET("/:hash", h.getOriginalURL)

	api := router.Group("/api")
	api.POST("/create", h.createShortURL)

	return router
}
