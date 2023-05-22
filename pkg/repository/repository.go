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
	Create(userId string, list todo.TodoList) (string, error)
	GetAll(userId string) ([]todo.TodoList, error)
	GetById(userId, listId string) (todo.TodoList, error)
	Delete(userId, listId string) error
	Update(userId, listId string, input todo.UpdateListInput) error
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
		TodoList:      NewTodoListMongo(db.Collection("todo_lists")),
		TodoItem:      nil,
	}
}
