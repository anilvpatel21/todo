package repository

import (
	"errors"
	"sync"

	"example.com/todo-app/internal/todo"
)

// InMemoryTodoRepository is an in-memory implementation of the TodoRepository interface.
type InMemoryTodoRepository struct {
	mu    sync.RWMutex
	todos map[string][]todo.Todo // Keyed by userID
}

// NewInMemoryTodoRepository creates a new instance of InMemoryTodoRepository.
func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos: make(map[string][]todo.Todo),
	}
}

// Create adds a new todo.Todo to the repository.
func (r *InMemoryTodoRepository) Create(todo todo.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Add the new todo.Todo to the user's list of Todos
	r.todos[todo.UserID] = append(r.todos[todo.UserID], todo)
	return nil
}

// GetByID retrieves a todo.Todo by ID for a specific user.
func (r *InMemoryTodoRepository) GetByID(id string, userID string) (*todo.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todos, exists := r.todos[userID]
	if !exists {
		return nil, errors.New("user not found")
	}

	for _, todo := range todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, errors.New("todo not found")
}

// GetByUserID retrieves all Todos for a specific user.
func (r *InMemoryTodoRepository) GetAllByUserID(userID string) ([]todo.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todos, exists := r.todos[userID]
	if !exists {
		return nil, errors.New("user not found")
	}
	return todos, nil
}

// Update modifies an existing todo.Todo in the repository.
func (r *InMemoryTodoRepository) Update(todo todo.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	todos, exists := r.todos[todo.UserID]
	if !exists {
		return errors.New("user not found")
	}

	// Find the todo.Todo and update it
	for i, existingTodo := range todos {
		if existingTodo.ID == todo.ID {
			todos[i] = todo
			r.todos[todo.UserID] = todos
			return nil
		}
	}
	return errors.New("todo not found")
}

// Delete removes a todo.Todo by ID for a specific user.
func (r *InMemoryTodoRepository) Delete(id string, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	todos, exists := r.todos[userID]
	if !exists {
		return errors.New("user not found")
	}

	// Find the todo.Todo and remove it
	for i, existingTodo := range todos {
		if existingTodo.ID == id {
			r.todos[userID] = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
