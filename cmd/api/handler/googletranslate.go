package handler

import (
	"api-google-translate/internal/googletranslate"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type (
	TranslateEntity struct {
		Text   string `json:"text" binding:"required"`
		Target string `json:"target" binding:"required"`
		Source string `json:"source"`
	}

	DetectEntity struct {
		Text string `json:"text" binding:"required"`
	}

	MotivationEntity struct {
		Motivation string `json:"motivation" binding:"required"`
	}
)

type GoogleTranslate struct {
	service googletranslate.Service
}

func NewGoogleTranslate(service googletranslate.Service) *GoogleTranslate {
	return &GoogleTranslate{
		service: service,
	}
}

func (handler *GoogleTranslate) GetWordTranslated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var translateEntity TranslateEntity

		if err := binding.JSON.Bind(ctx.Request, &translateEntity); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		textTranslated, err := handler.service.GetWordTranslated(ctx, translateEntity.Text, translateEntity.Target, translateEntity.Source)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, textTranslated)
	}
}

//TODO: Verify add cache
func (handler *GoogleTranslate) GetSupportedLanguages() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		languages, err := handler.service.GetSupportedLanguages(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, languages)
	}
}

func (handler *GoogleTranslate) GetDetectedLanguage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var detectLanguage DetectEntity
		if err := binding.JSON.Bind(ctx.Request, &detectLanguage); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		detectedLanguage, err := handler.service.GetDetectedLanguage(ctx, detectLanguage.Text)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, detectedLanguage)
	}
}

func (handler *GoogleTranslate) GetMotivationTranslated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		motivationEntity := &MotivationEntity{}
		if err := binding.JSON.Bind(ctx.Request, &motivationEntity); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		motivationTranslated, err := handler.service.GetMotivationTranslated(ctx, motivationEntity.Motivation)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, motivationTranslated)
	}
}
