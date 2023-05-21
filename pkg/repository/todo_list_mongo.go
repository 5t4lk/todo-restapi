package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *TodoListMongo) GetAll(userId string) ([]todo.TodoList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": userId}

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var lists []todo.TodoList
	for cur.Next(ctx) {
		var list todo.TodoList
		if err := cur.Decode(&list); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return lists, nil
}
