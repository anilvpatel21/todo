package todo

import "time"

// Attachment represents an uploaded file (photo, video, audio, etc.).
type Attachment struct {
	// ID: A unique identifier for the attachment, used to distinguish between different attachments.
	ID string `json:"id"`

	// FileName: The name of the attached file (e.g., `image.jpg`, `document.pdf`).
	FileName string `json:"file_name"`

	// FileType: The MIME type of the attachment (e.g., `image/jpeg`, `audio/mp3`, `video/mp4`).
	FileType string `json:"file_type"`

	// FilePath: The file path where the attachment is stored (e.g., `/uploads/images/image.jpg`).
	FilePath string `json:"file_path"`
}

// Todo represents a single todo item with attachments.
type Todo struct {
	// ID: A unique identifier for the todo item, used to distinguish between different todos.
	ID string `json:"id"`

	// UserID: The identifier for the user who created the todo, ensuring that todos are tied to specific users.
	UserID string `json:"user_id"`

	// Title: A short, optional string that provides a summary or name for the todo item.
	Title string `json:"title,omitempty"`

	// Description: A longer, optional text that describes the task or objective of the todo item.
	Description string `json:"description,omitempty"`

	// CreatedAt: The timestamp when the todo item was created.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt: The timestamp when the todo item was last updated.
	UpdatedAt time.Time `json:"updated_at"`

	// Attachments: A list of attachments (files) associated with the todo item.
	Attachments []Attachment `json:"attachments"`
}

// TodoRepository defines methods for interacting with the data store for todo items.
type TodoRepository interface {
	// Create adds a new todo item to the repository (e.g., database).
	Create(todo Todo) error

	// GetByID retrieves a todo item by its ID for a specific user.
	GetByID(id string, userID string) (*Todo, error)

	// GetAllByUserID retrieves all todo items for a specific user.
	GetAllByUserID(userID string) ([]Todo, error)

	// Update modifies an existing todo item with new data.
	Update(todo Todo) error

	// Delete removes a todo item by its ID for a specific user.
	Delete(id string, userID string) error
}
