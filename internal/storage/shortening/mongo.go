package shortening

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"url-shotener-api/internal/models"
)

type mgo struct {
	db *mongo.Database
}

func NewMongoDB(client *mongo.Client) *mgo {
	return &mgo{db: client.Database("url-shortener")}
}

func (m *mgo) col() *mongo.Collection {
	return m.db.Collection("shortenings")
}

func (m *mgo) Put(ctx context.Context, shortening models.Shortening) (*models.Shortening, error) {
	const op = "shortening.mgo.Put"

	shortening.CreateAt = time.Now().UTC()

	count, err := m.col().CountDocuments(ctx, bson.M{"identifier": shortening.Identifier})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if count > 0 {
		return nil, fmt.Errorf("%s: %w", op, models.ErrIdentifierExists)
	}

	_, err = m.db.Collection("shortenings").InsertOne(ctx, mgoShorteningFromModel(shortening))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &shortening, nil
}

type mgoShortening struct {
	Identifier  string    `bson:"_id"`
	OriginalURL string    `bson:"original_url"`
	Visits      int64     `bson:"visits"`
	CreateAt    time.Time `bson:"create_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

func mgoShorteningFromModel(shortening models.Shortening) mgoShortening {
	return mgoShortening{
		Identifier:  shortening.Identifier,
		OriginalURL: shortening.OriginalURL,
		Visits:      shortening.Visits,
		CreateAt:    shortening.CreateAt,
		UpdatedAt:   shortening.UpdatedAt,
	}
}

func modelShorteningFromMgo(shortening mgoShortening) *models.Shortening {
	return &models.Shortening{
		Identifier:  shortening.Identifier,
		OriginalURL: shortening.OriginalURL,
		Visits:      shortening.Visits,
		CreateAt:    shortening.CreateAt,
		UpdatedAt:   shortening.UpdatedAt,
	}
}
