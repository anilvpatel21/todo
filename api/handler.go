package api

import (
	"example.com/todo-app/internal/todo"
)

// TodoHandler handles HTTP requests related to TODOs.
type TodoHandler struct {
	service todo.TodoService
}

// NewTodoHandler creates a new instance of TodoHandler.
func NewTodoHandler(service todo.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}
