package storage

import (
	"sync"
	"github.com/example/todo-service/internal/domain"
	"github.com/example/todo-service/internal/port"
)

type inMemoryRepository struct {
	mu    sync.RWMutex
	todos map[int]*domain.Todo
	id    int
}

func NewInMemoryRepository() port.TodoRepository {
	return &inMemoryRepository{
		todos: make(map[int]*domain.Todo),
		id:    0,
	}
}

func (r *inMemoryRepository) Create(title string) *domain.Todo {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.id++
	todo := &domain.Todo{
		ID:        r.id,
		Title:     title,
		Completed: false,
	}
	r.todos[r.id] = todo
	return todo
}

func (r *inMemoryRepository) GetAll() []*domain.Todo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todos := make([]*domain.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}
	return todos
}

func (r *inMemoryRepository) GetByID(id int) *domain.Todo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.todos[id]
}

func (r *inMemoryRepository) Update(id int, title string, completed bool) *domain.Todo {
	r.mu.Lock()
	defer r.mu.Unlock()

	if todo, exists := r.todos[id]; exists {
		todo.Title = title
		todo.Completed = completed
		return todo
	}
	return nil
}

func (r *inMemoryRepository) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.todos[id]; exists {
		delete(r.todos, id)
		return true
	}
	return false
}
