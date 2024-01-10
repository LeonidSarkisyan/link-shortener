package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"url-shotener-api/internal/models"
)

type AuthMongo struct {
	db *mongo.Database
}

func newAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{db}
}

func (a *AuthMongo) Create(userId uint, user models.UserInput) (uint, error) {
	userWithId := models.User{
		Id:       userId,
		Username: user.Username,
		Password: user.Password,
	}
	userExist, err := a.GetByUsername(user.Username)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return 0, err
	}
	if userExist.Id != 0 {
		return 0, UserAlreadyExists
	}
	_, err = a.db.Collection(userCollection).InsertOne(context.TODO(), userWithId)
	if err != nil {
		return 0, err
	}
	return userWithId.Id, err
}

func (a *AuthMongo) GetByUsername(username string) (models.User, error) {
	var user models.User
	err := a.db.Collection(userCollection).FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
