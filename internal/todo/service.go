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

// CreateTodo creates a new TODO item for a specific user.
func (s *Service) CreateTodo(userID, title, description string) (*Todo, error) {
	return nil, nil
}

// GetTodosByUserID returns all TODO items for a specific user.
func (s *Service) GetTodosByID(id, userID string) (*Todo, error) {
	return nil, nil
}

// GetTodosByUserID returns all TODO items for a specific user.
func (s *Service) GetTodosByUserID(userID string) ([]Todo, error) {
	return nil, nil
}

// UpdateTodo updates an existing TODO item for a specific user.
func (s *Service) UpdateTodo(id, userID, title, description string) (*Todo, error) {
	return nil, nil
}

// DeleteTodo deletes a TODO item for a specific user.
func (s *Service) DeleteTodo(id, userID string) error {
	return nil
}

// AddAttachment adds an attachment to a TODO item for a specific user.
func (s *Service) AddAttachment(todoID, userID, fileName, fileType, filePath string) (*Attachment, error) {
	return nil, nil
}

// DeleteAttachment removes an attachment from a TODO item for a specific user.
func (s *Service) DeleteAttachment(todoID, userID, attachmentID string) (*Todo, error) {
	return nil, nil
}
