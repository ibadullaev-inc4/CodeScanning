package port

import "github.com/example/todo-service/internal/domain"

// TodoRepository - интерфейс для работы с хранилищем
type TodoRepository interface {
	Create(title string) *domain.Todo
	GetAll() []*domain.Todo
	GetByID(id int) *domain.Todo
	Update(id int, title string, completed bool) *domain.Todo
	Delete(id int) bool
}

// TodoService - интерфейс для бизнес-логики
type TodoService interface {
	CreateTodo(title string) *domain.Todo
	GetAllTodos() []*domain.Todo
	GetTodo(id int) *domain.Todo
	UpdateTodo(id int, title string, completed bool) *domain.Todo
	DeleteTodo(id int) bool
}
