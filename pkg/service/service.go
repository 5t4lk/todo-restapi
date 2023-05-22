package service

import (
	"todo"
	"todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (string, error)
	GenerateToken(username, password string) (string, error)
	Validate(token string) (string, error)
}

type TodoList interface {
	Create(userId string, list todo.TodoList) (string, error)
	GetAll(userId string) ([]todo.TodoList, error)
	GetById(userId, listId string) (todo.TodoList, error)
	Delete(userId, listId string) error
	Update(userId, listId string, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId string, item todo.TodoItem) (string, error)
	GetAll(userId, listId string) ([]todo.TodoItem, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
