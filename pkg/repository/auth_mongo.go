package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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
	user.Id = primitive.NewObjectID().Hex()

	result, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *AuthMongo) GetUser(username, password string) (todo.User, error) {
	var user todo.User

	filter := bson.M{
		"username": username,
		"password": password,
	}

	if err := r.collection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return todo.User{}, err
	}

	return user, nil
}
