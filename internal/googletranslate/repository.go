package googletranslate

import (
	"api-google-translate/internal/domain"
	"api-google-translate/internal/googletranslate/http"
)

type Repository interface {
	GetWordTranslated(text, target, source string) (domain.GoogleTranslate, error)
	GetSupportedLanguages() ([]domain.GoogleLanguage, error)
	GetDetectedLanguage(text string) (domain.GoogleDetector, error)
}

func NewRepository() Repository {
	return http.NewHttpRepository().(Repository)
}
