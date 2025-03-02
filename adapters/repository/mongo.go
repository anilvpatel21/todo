package repository

import (
	"context"
	"errors"

	"example.com/todo-app/internal/todo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTodoRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func Connect() (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://anilpatelyt:<db_password>@cluster0.0u9qr.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// NewMongoTodoRepository creates a new instance of MongoTodoRepository.
func NewMongoTodoRepository(client *mongo.Client, dbName, collectionName string) *MongoTodoRepository {
	return &MongoTodoRepository{
		client:     client,
		collection: client.Database(dbName).Collection(collectionName),
	}
}

// Create adds a new todo.Todo to the MongoDB repository.
func (r *MongoTodoRepository) Create(todo todo.Todo) error {
	_, err := r.collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a todo.Todo by ID for a specific user from MongoDB.
func (r *MongoTodoRepository) GetByID(id string, userid string) (*todo.Todo, error) {
	filter := bson.M{"id": id, "userid": userid}
	var result todo.Todo
	err := r.collection.FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("todo not found")
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByUserID retrieves all Todos for a specific user from MongoDB.
func (r *MongoTodoRepository) GetAllByUserID(userid string) ([]todo.Todo, error) {
	filter := bson.M{"userid": userid}
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var todos []todo.Todo
	for cursor.Next(context.Background()) {
		var todo todo.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

// Update modifies an existing todo.Todo in MongoDB.
func (r *MongoTodoRepository) Update(todo todo.Todo) error {
	filter := bson.M{"id": todo.ID, "userid": todo.UserID}
	update := bson.M{"$set": todo}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a todo.Todo by ID for a specific user from MongoDB.
func (r *MongoTodoRepository) Delete(id string, userid string) error {
	filter := bson.M{"id": id, "userid": userid}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
