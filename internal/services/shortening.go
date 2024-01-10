package services

import (
	"github.com/google/uuid"
	"url-shotener-api/internal/models"
	"url-shotener-api/internal/repositories"
	"url-shotener-api/internal/shorten"
)

type ShorteningService struct {
	repository *repositories.Repository
}

func NewShorteningService(repo *repositories.Repository) *ShorteningService {
	return &ShorteningService{repository: repo}
}

func (s ShorteningService) Create(shortening models.ShortenInput) (string, error) {
	if shortening.Identifier == nil {
		id := shorten.Shorten(uuid.New().ID())
		shortening.Identifier = &id
	}
	return s.repository.Shortening.Create(shortening)
}

func (s ShorteningService) Get(alias string) (string, error) {
	return s.repository.Shortening.Get(alias)
}
