package main

import (
	"context"
	"fmt"
	"log"

	"example.com/todo-app/adapters/repository"
	"example.com/todo-app/api"
	"example.com/todo-app/config"
	internal "example.com/todo-app/internal/todo"
)

func main() {

	err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	client, err := repository.Connect()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	repo := repository.NewMongoTodoRepository(client, "todoAppDB", "todos")
	// Initialize the service with the repository
	service := internal.NewService(repo)

	router := api.SetupRoutes(service)

	// Start the server
	port := "8080"
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(router.Run(":" + port))
}
