package todo

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
