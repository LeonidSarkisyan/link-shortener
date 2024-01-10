package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"url-shotener-api/internal/models"
)

type Auth interface {
	Create(userId uint, user models.UserInput) (uint, error)
	GetByUsername(username string) (models.User, error)
}

type Shortening interface {
	Create(shortening models.ShortenInput) (string, error)
	Get(id string) (string, error)
}

type Repository struct {
	Auth
	Shortening
}

func NewRepositoryFromMangoDB(db *mongo.Database) *Repository {
	return &Repository{
		Shortening: newShorteningMongo(db),
		Auth:       newAuthMongo(db),
	}
}
