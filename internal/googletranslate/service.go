package googletranslate

import (
	"api-google-translate/internal/domain"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetWordTranslated(ctx context.Context, text, target, source string) (domain.GoogleTranslate, error)
	GetSupportedLanguages(ctx context.Context) ([]domain.GoogleLanguage, error)
	GetDetectedLanguage(ctx context.Context, text string) (domain.GoogleDetector, error)
	GetMotivationTranslated(ctx *gin.Context, motivation string) (domain.Motivation, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

const (
	LANGUAGE_DETECTED_IS_NOT_RELIABLE = "The language detected is not reliable"
)

func (s *service) GetWordTranslated(ctx context.Context, text, target, source string) (domain.GoogleTranslate, error) {
	textTranslated, err := s.repository.GetWordTranslated(text, target, source)
	if err != nil {
		return domain.GoogleTranslate{}, err
	}

	return textTranslated, nil
}

func (s *service) GetSupportedLanguages(ctx context.Context) ([]domain.GoogleLanguage, error) {
	languages, err := s.repository.GetSupportedLanguages()
	if err != nil {
		return []domain.GoogleLanguage{}, err
	}

	return languages, nil
}

func (s *service) GetDetectedLanguage(ctx context.Context, text string) (domain.GoogleDetector, error) {
	detectedLanguage, err := s.repository.GetDetectedLanguage(text)
	if err != nil {
		return domain.GoogleDetector{}, err
	}

	return detectedLanguage, nil
}

func (s *service) GetMotivationTranslated(ctx *gin.Context, motivation string) (domain.Motivation, error) {
	detectedLanguage, err := s.GetDetectedLanguage(ctx, motivation)
	if err != nil {
		return domain.Motivation{}, err
	}

	if detectedLanguage.IsReliable {
		return domain.Motivation{}, errors.New(LANGUAGE_DETECTED_IS_NOT_RELIABLE)
	}

	textTranslated, err := s.GetWordTranslated(ctx, motivation, "es", detectedLanguage.Language)
	if err != nil {
		return domain.Motivation{}, err
	}

	return domain.Motivation{Body: textTranslated.Text}, nil
}
