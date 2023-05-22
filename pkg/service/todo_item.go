package service

import (
	"todo"
	"todo/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId string, item todo.TodoItem) (string, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return "", err // list does not exist or not belong to user
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId string) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}
