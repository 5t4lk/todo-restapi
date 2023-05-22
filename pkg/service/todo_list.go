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

func (s *TodoListService) GetAll(userId string) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId string) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}
