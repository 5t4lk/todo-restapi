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
}

type TodoItem interface {
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
	}
}
