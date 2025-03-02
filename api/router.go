package api

import (
	"example.com/todo-app/internal/todo"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(service todo.TodoService) *gin.Engine {
	handler := NewTodoHandler(service)
	// Set up the Gin router
	r := gin.Default()

	// Define the routes
	r.POST("/todos", handler.CreateTodoHandler)
	r.GET("/todos", handler.GetAllTodosHandler)
	r.GET("/todos/:todoID", handler.GetTodosHandler)
	r.PUT("/todos/:todoID", handler.UpdateTodoHandler)
	r.DELETE("/todos/:todoID", handler.DeleteTodoHandler)
	r.POST("/todos/:todoID/attachments", handler.UploadAttachmentHandler)
	//r.PUT("/todos/:todoID/attachments", handler.AddAttachmentHandler)
	r.DELETE("/todos/:todoID/attachments/:attachmentID", handler.DeleteAttachmentHandler)

	return r
}
