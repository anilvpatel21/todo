package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"example.com/todo-app/internal/todo"
	"github.com/gin-gonic/gin"
)

// TodoHandler handles HTTP requests related to TODOs.
type TodoHandler struct {
	service todo.TodoService
}

// NewTodoHandler creates a new instance of TodoHandler.
func NewTodoHandler(service todo.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func authenticate(c *gin.Context) (string, error) {
	// For demo purposes, we'll assume a mock user ID
	// Replace this with your actual authentication logic (JWT, session, etc.)
	return "mock-user-id", nil
}

// CreateTodoHandler handles the creation of a new TODO item for a user.
func (h *TodoHandler) CreateTodoHandler(c *gin.Context) {
	userID, err := authenticate(c) // Authentication logic
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	var request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	todo, err := h.service.CreateTodo(userID, request.Title, request.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// GetTodosHandler returns all TODO items for the authenticated user.
func (h *TodoHandler) GetTodosHandler(c *gin.Context) {
	userID, err := authenticate(c) // Authentication logic
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Get the todoID from the URL parameter
	todoID := c.Param("todoID")

	todo, err := h.service.GetTodosByID(todoID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// GetTodosHandler returns all TODO items for the authenticated user.
func (h *TodoHandler) GetAllTodosHandler(c *gin.Context) {
	userID, err := authenticate(c) // Authentication logic
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	todos, err := h.service.GetTodosByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// UpdateTodoHandler updates an existing TODO item for a user.
func (h *TodoHandler) UpdateTodoHandler(c *gin.Context) {
	userID, err := authenticate(c) // Authentication logic
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Get the todoID from the URL parameter
	todoID := c.Param("todoID")

	var request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	todo, err := h.service.UpdateTodo(todoID, userID, request.Title, request.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodoHandler deletes a TODO item for a user.
func (h *TodoHandler) DeleteTodoHandler(c *gin.Context) {
	userID, err := authenticate(c) // Authentication logic
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	todoID := c.Param("todoID")

	err = h.service.DeleteTodo(todoID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.Status(http.StatusNoContent)
}

// UploadAttachmentHandler handles file uploads and saves the file locally.
func (h *TodoHandler) UploadAttachmentHandler(c *gin.Context) {
	// Get userID from authentication (implement this logic as needed)
	userID, err := authenticate(c) // Authentication logic
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Get todoID from URL parameter
	todoID := c.Param("todoID")

	// Retrieve the file from the form data (key: "file")
	file, _ := c.FormFile("file")
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Define the file path to store the file
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	// Generate a unique file name to prevent name collisions (you can also use the original file name)
	fileName := fmt.Sprintf("%s-%d-%s", time.Now().Format("20060102150405"), time.Now().UnixNano(), file.Filename)
	filePath := filepath.Join(uploadDir, fileName)

	// Save the file to the uploads directory
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Create attachment information
	reqAttachment := todo.Attachment{
		FileName: fileName,
		FileType: file.Header.Get("Content-Type"),
		FilePath: filePath,
	}

	// Call service to add this attachment to the TODO item
	uploadedAttachment, err := h.service.AddAttachment(todoID, userID, reqAttachment.FileName, reqAttachment.FileType, reqAttachment.FilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add attachment to todo"})
		return
	}

	// Return file information in the response
	c.JSON(http.StatusOK, gin.H{
		"id":       uploadedAttachment.ID,
		"fileName": uploadedAttachment.FileName,
		"fileType": uploadedAttachment.FileType,
		"filePath": uploadedAttachment.FilePath,
	})
}

// DeleteAttachmentHandler removes an attachment from a TODO item.
func (h *TodoHandler) DeleteAttachmentHandler(c *gin.Context) {
	userID, err := authenticate(c) // Authentication logic
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	todoID := c.Param("todoID")
	attachmentID := c.Param("attachmentID")

	todo, err := h.service.DeleteAttachment(todoID, userID, attachmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete attachment"})
		return
	}

	c.JSON(http.StatusOK, todo)
}
