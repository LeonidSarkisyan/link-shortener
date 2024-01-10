package storages

import "go.mongodb.org/mongo-driver/mongo"

type MangoStorage struct {
	Db *mongo.Database
}

func NewMongoStorage(client *mongo.Client, databaseName string) *MangoStorage {
	return &MangoStorage{Db: client.Database(databaseName)}
}

func (s *MangoStorage) Collection(name string) *mongo.Collection {
	return s.Db.Collection(name)
}
