package handler

import (
	"api-google-translate/internal/motivation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service motivation.Service
}

func NewMotivation(service motivation.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (handler *Handler) GetMotivation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		motivationObtained, err := handler.service.GetMotivation()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, motivationObtained)
	}
}
