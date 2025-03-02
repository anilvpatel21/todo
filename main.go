package main

import (
	"fmt"
	"log"

	"example.com/todo-app/adapters/repository"
	"example.com/todo-app/api"
	internal "example.com/todo-app/internal/todo"
)

func main() {
	repo := repository.NewInMemoryTodoRepository()
	// Initialize the service with the repository
	service := internal.NewService(repo)

	router := api.SetupRoutes(service)

	// Start the server
	port := "8080"
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(router.Run(":" + port))
}
