package handler

import (
	"github.com/fidesy/ozon-test/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.GET("/:hash", h.getOriginalURL)

	api := router.Group("/api")
	api.POST("/create", h.createShortURL)

	return router
}
