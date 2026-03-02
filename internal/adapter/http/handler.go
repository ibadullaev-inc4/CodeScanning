package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/example/todo-service/internal/port"
)

type todoHandler struct {
	service port.TodoService
}

func NewTodoHandler(service port.TodoService) *todoHandler {
	return &todoHandler{service: service}
}

func (h *todoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todos := h.service.GetAllTodos()
	json.NewEncoder(w).Encode(todos)
}

func (h *todoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo := h.service.GetTodo(id)
	if todo == nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *todoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	todo := h.service.CreateTodo(req.Title)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (h *todoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	todo := h.service.UpdateTodo(id, req.Title, req.Completed)
	if todo == nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *todoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if !h.service.DeleteTodo(id) {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *todoHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func RegisterRoutes(router *mux.Router, service port.TodoService) {
	handler := NewTodoHandler(service)

	router.HandleFunc("/health", handler.Health).Methods("GET")
	router.HandleFunc("/todos", handler.GetAllTodos).Methods("GET")
	router.HandleFunc("/todos", handler.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", handler.GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", handler.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", handler.DeleteTodo).Methods("DELETE")
}
