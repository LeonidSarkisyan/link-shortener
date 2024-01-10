package shortening

import (
	"context"
	"sync"
	"time"
	"url-shotener-api/internal/models"
)

type inMemory struct {
	m sync.Map
}

func NewInMemory() *inMemory {
	return &inMemory{}
}

func (s *inMemory) Put(_ context.Context, shortening models.Shortening) (*models.Shortening, error) {
	if _, exists := s.m.Load(shortening.Identifier); exists {
		return nil, models.ErrIdentifierExists
	}

	shortening.CreateAt = time.Now().UTC()

	s.m.Store(shortening.Identifier, shortening)

	return &shortening, nil
}

func (s *inMemory) Get(_ context.Context, identifier string) (*models.Shortening, error) {
	v, ok := s.m.Load(identifier)
	if !ok {
		return nil, models.ErrNotFound
	}

	shortening := v.(models.Shortening)

	return &shortening, nil
}

func (s *inMemory) IncrementVisits(_ context.Context, identifier string) error {
	v, ok := s.m.Load(identifier)
	if !ok {
		return models.ErrNotFound
	}

	shortening := v.(models.Shortening)
	shortening.Visits++

	s.m.Store(identifier, shortening)

	return nil
}
