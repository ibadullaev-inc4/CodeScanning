package service

import (
	"github.com/example/todo-service/internal/domain"
	"github.com/example/todo-service/internal/port"
)

type todoService struct {
	repo port.TodoRepository
}

func NewTodoService(repo port.TodoRepository) port.TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) CreateTodo(title string) *domain.Todo {
	return s.repo.Create(title)
}

func (s *todoService) GetAllTodos() []*domain.Todo {
	return s.repo.GetAll()
}

func (s *todoService) GetTodo(id int) *domain.Todo {
	return s.repo.GetByID(id)
}

func (s *todoService) UpdateTodo(id int, title string, completed bool) *domain.Todo {
	return s.repo.Update(id, title, completed)
}

func (s *todoService) DeleteTodo(id int) bool {
	return s.repo.Delete(id)
}
