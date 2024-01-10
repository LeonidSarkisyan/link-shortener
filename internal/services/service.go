package services

import (
	"url-shotener-api/internal/models"
	"url-shotener-api/internal/repositories"
)

type Auth interface {
	Register(user models.UserInput) (uint, error)
	Login(user models.UserInput) (string, error)
}

type Shortening interface {
	Create(shortening models.ShortenInput) (string, error)
	Get(alias string) (string, error)
}

type Service struct {
	Auth
	Shortening
}

func NewServices(repository *repositories.Repository) *Service {
	return &Service{
		Shortening: NewShorteningService(repository),
		Auth:       NewAuthService(repository),
	}
}
