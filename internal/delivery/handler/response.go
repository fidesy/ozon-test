package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, err error) {
	log.Println(err.Error())
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: err.Error()})
}