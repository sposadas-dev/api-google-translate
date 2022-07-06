package motivation

import "api-google-translate/internal/motivation/http"

type Repository interface {
	GetMotivation() (string, error)
}

func NewRepository() Repository {
	return http.NewHttpRepository().(Repository)
}
