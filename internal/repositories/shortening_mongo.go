package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"url-shotener-api/internal/models"
)

type ShorteningMongo struct {
	db *mongo.Database
}

func newShorteningMongo(db *mongo.Database) *ShorteningMongo {
	return &ShorteningMongo{db: db}
}

func (s ShorteningMongo) Create(shortening models.ShortenInput) (string, error) {
	_, err := s.db.Collection(shorteningCollection).InsertOne(context.TODO(), shortening)
	if err != nil {
		return "", err
	}
	return *shortening.Identifier, nil
}

func (s ShorteningMongo) Get(alias string) (string, error) {
	var result models.ShortenInput
	filter := bson.D{{"identifier", alias}}
	err := s.db.Collection(shorteningCollection).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.RawURL, nil
}
