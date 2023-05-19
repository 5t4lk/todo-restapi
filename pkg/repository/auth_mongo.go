package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"todo"
)

type AuthMongo struct {
	collection *mongo.Collection
}

func NewAuthMongo(collection *mongo.Collection) *AuthMongo {
	return &AuthMongo{collection: collection}
}

func (r *AuthMongo) CreateUser(user todo.User) (string, error) {
	result, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}
