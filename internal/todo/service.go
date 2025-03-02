package todo

import (
	"errors"
	"fmt"
	"time"

	"example.com/todo-app/internal/helper"
)

// Service encapsulates the business logic.
type Service struct {
	repo TodoRepository
}

// TodoService defines the service layer for managing Todo items.
type TodoService interface {
	CreateTodo(userID, title, description string) (*Todo, error)
	GetTodosByID(id, userID string) (*Todo, error)
	GetTodosByUserID(userID string) ([]Todo, error)
	UpdateTodo(id, userID, title, description string) (*Todo, error)
	DeleteTodo(id, userID string) error
	AddAttachment(todoID, userID, fileName, fileType, filePath string) (*Attachment, error)
	DeleteAttachment(todoID, userID, attachmentID string) (*Todo, error)
}

// NewService creates a new Todo Service.
func NewService(repo TodoRepository) *Service {
	return &Service{repo: repo}
}

// CreateTodo creates a new TODO item for a specific user.
func (s *Service) CreateTodo(userID, title, description string) (*Todo, error) {
	todo := Todo{
		ID:          helper.GenerateID(),
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UserID:      userID, // Associate the TODO item with the user
		Attachments: []Attachment{},
	}

	err := s.repo.Create(todo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &todo, nil
}

// GetTodosByUserID returns all TODO items for a specific user.
func (s *Service) GetTodosByID(id, userID string) (*Todo, error) {
	return s.repo.GetByID(id, userID)
}

// GetTodosByUserID returns all TODO items for a specific user.
func (s *Service) GetTodosByUserID(userID string) ([]Todo, error) {
	return s.repo.GetAllByUserID(userID)
}

// UpdateTodo updates an existing TODO item for a specific user.
func (s *Service) UpdateTodo(id, userID, title, description string) (*Todo, error) {
	todo, err := s.repo.GetByID(id, userID)
	if err != nil {
		return nil, err
	}

	todo.Title = title
	todo.Description = description
	todo.UpdatedAt = time.Now()

	err = s.repo.Update(*todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// DeleteTodo deletes a TODO item for a specific user.
func (s *Service) DeleteTodo(id, userID string) error {
	return s.repo.Delete(id, userID)
}

// AddAttachment adds an attachment to a TODO item for a specific user.
func (s *Service) AddAttachment(todoID, userID, fileName, fileType, filePath string) (*Attachment, error) {
	todo, err := s.repo.GetByID(todoID, userID)
	if err != nil {
		return nil, err
	}

	attachment := Attachment{
		ID:       helper.GenerateID(),
		FileName: fileName,
		FileType: fileType,
		FilePath: filePath,
	}

	todo.Attachments = append(todo.Attachments, attachment)
	todo.UpdatedAt = time.Now()

	err = s.repo.Update(*todo)
	if err != nil {
		return nil, err
	}

	return &attachment, nil
}

// DeleteAttachment removes an attachment from a TODO item for a specific user.
func (s *Service) DeleteAttachment(todoID, userID, attachmentID string) (*Todo, error) {
	todo, err := s.repo.GetByID(todoID, userID)
	if err != nil {
		return nil, err
	}

	// Find and remove the attachment by ID
	for i, att := range todo.Attachments {
		if att.ID == attachmentID {
			todo.Attachments = append(todo.Attachments[:i], todo.Attachments[i+1:]...)
			todo.UpdatedAt = time.Now()
			err = s.repo.Update(*todo)
			if err != nil {
				return nil, err
			}
			return todo, nil
		}
	}
	return nil, errors.New("attachment not found")
}
