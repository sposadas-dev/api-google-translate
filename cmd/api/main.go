package main

import (
	"api-google-translate/cmd/api/handler"
	"api-google-translate/internal/googletranslate"
	"api-google-translate/internal/motivation"
	"api-google-translate/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

const (
	API_VERSION          = "/api/v1"
	GOOGLE_TRANSLATE_API = "/google/translate"
	MOTIVATIONS_API      = "/motivations"
)

type HealthChecker struct{}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	server := gin.New()
	checker := HealthChecker{}

	mapRoutes(server, checker)
	return server.Run()
}

func mapRoutes(server *gin.Engine, checker HealthChecker) {
	server.GET("/ping", checker.HealthCheck)

	googleTranslateRepository := googletranslate.NewRepository()
	googleTranslateService := googletranslate.NewService(googleTranslateRepository)
	googleTranslateHandler := handler.NewGoogleTranslate(googleTranslateService)

	googleTranslateAPIGroup := server.Group(API_VERSION + GOOGLE_TRANSLATE_API)

	{
		googleTranslateAPIGroup.POST("/", web.TokenAuthMiddleware(), googleTranslateHandler.GetWordTranslated())
		googleTranslateAPIGroup.GET("/languages", web.TokenAuthMiddleware(), googleTranslateHandler.GetSupportedLanguages())
		googleTranslateAPIGroup.POST("/detect", web.TokenAuthMiddleware(), googleTranslateHandler.GetDetectedLanguage())
		googleTranslateAPIGroup.POST("/motivation", web.TokenAuthMiddleware(), googleTranslateHandler.GetMotivationTranslated())
	}

	motivationsRepository := motivation.NewRepository()
	motivationsService := motivation.NewService(motivationsRepository)
	motivationsHandler := handler.NewMotivation(motivationsService)

	motivationsAPIGroup := server.Group(API_VERSION + MOTIVATIONS_API)

	{
		motivationsAPIGroup.GET("/", web.TokenAuthMiddleware(), motivationsHandler.GetMotivation())
	}
}

func (h HealthChecker) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}
