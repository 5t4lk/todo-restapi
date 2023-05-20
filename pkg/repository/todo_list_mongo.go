package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"todo"
)

type TodoListMongo struct {
	collection *mongo.Collection
}

func NewTodoListMongo(collection *mongo.Collection) *TodoListMongo {
	return &TodoListMongo{collection: collection}
}

func (r *TodoListMongo) Create(userId string, list todo.TodoList) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	list.Id = userId

	result, err := r.collection.InsertOne(ctx, list)
	if err != nil {
		return "", err
	}

	listID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("failed to get inserted ID")
	}

	return listID.Hex(), nil
}
