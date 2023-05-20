package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"todo"
)

type Authorization interface {
	CreateUser(user todo.User) (string, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(client *mongo.Client, dbName string) *Repository {
	db := client.Database(dbName)

	return &Repository{
		Authorization: NewAuthMongo(db.Collection("users")),
		TodoList:      nil,
		TodoItem:      nil,
	}
}
