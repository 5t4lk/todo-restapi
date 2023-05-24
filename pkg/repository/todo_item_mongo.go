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

type todoItemMongo struct {
	collection *mongo.Collection
}

func NewTodoItemMongo(collection *mongo.Collection) *todoItemMongo {
	return &todoItemMongo{collection: collection}
}

func (r *todoItemMongo) Create(listId string, item todo.TodoItem) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := map[string]interface{}{
		"title":       item.Title,
		"description": item.Description,
		"list_id":     listId,
	}

	result, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}

	itemId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("failed to get inserted ID")
	}

	return itemId.Hex(), nil
}

func (r *todoItemMongo) GetAll(userId, listId string) ([]todo.TodoItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"list_id": listId,
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []todo.TodoItem

	for cursor.Next(ctx) {
		var item todo.TodoItem
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *todoItemMongo) GetById(userId, listId string) (todo.TodoItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return todo.TodoItem{}, err
	}

	filter := bson.M{
		"_id": _id,
	}

	var item todo.TodoItem
	err = r.collection.FindOne(ctx, filter).Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return todo.TodoItem{}, errors.New("list not found")
		}
		return todo.TodoItem{}, err
	}

	return item, nil
}

func (r *todoItemMongo) Delete(userId, listId string) error {
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

func (r *todoItemMongo) Update(userId, listId string, input todo.UpdateItemInput) error {
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
	if input.Done != nil {
		update["done"] = *input.Done
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
