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

	doc := map[string]interface{}{
		"title":       list.Title,
		"description": list.Description,
		"user_id":     userId,
	}

	result, err := r.collection.InsertOne(ctx, doc)
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

	filter := bson.M{"user_id": userId}

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

func (r *TodoListMongo) GetById(userId, listId string) (todo.TodoList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return todo.TodoList{}, err
	}

	filter := bson.M{
		"_id": _id,
	}

	var list todo.TodoList
	err = r.collection.FindOne(ctx, filter).Decode(&list)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return todo.TodoList{}, errors.New("list not found")
		}
		return todo.TodoList{}, err
	}

	return list, nil
}

func (r *TodoListMongo) Delete(userId, listId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": _id,
	}

	_, err = r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New("error occurred while deleting list")
	}

	return nil
}

func (r *TodoListMongo) Update(userId, listId string, input todo.UpdateListInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return err
	}

	update := map[string]interface{}{}

	if input.Title != "" {
		update["title"] = input.Title
	}
	if input.Description != "" {
		update["description"] = input.Description
	}

	filter := bson.M{
		"_id": _id,
	}

	updateQuery := bson.M{"$set": update}

	_, err = r.collection.UpdateOne(ctx, filter, updateQuery)
	if err != nil {
		return err
	}

	return nil
}
