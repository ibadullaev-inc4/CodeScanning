package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/example/todo-service/internal/adapter/http"
	"github.com/example/todo-service/internal/adapter/storage"
	"github.com/example/todo-service/internal/service"
)

func main() {
	// Adapter layer: инициализация хранилища
	repository := storage.NewInMemoryRepository()

	// Service layer: бизнес-логика
	todoService := service.NewTodoService(repository)

	// Adapter layer: HTTP обработчики
	router := mux.NewRouter()
	http.RegisterRoutes(router, todoService)

	// Запуск сервера
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
