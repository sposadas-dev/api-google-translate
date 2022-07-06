package motivation

import "api-google-translate/internal/domain"

type Service interface {
	GetMotivation() (domain.Motivation, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) GetMotivation() (domain.Motivation, error) {
	motivation, err := s.repository.GetMotivation()
	if err != nil {
		return domain.Motivation{}, err
	}

	return domain.Motivation{Body: motivation}, nil
}
