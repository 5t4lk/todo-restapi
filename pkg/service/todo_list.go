package service

import (
	"todo"
	"todo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId string, list todo.TodoList) (string, error) {
	return s.repo.Create(userId, list)
}
