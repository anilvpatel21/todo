package helper

import "github.com/google/uuid"

// Helper function to generate unique ID for Todo items (simplified).
func GenerateID() string {
	return uuid.NewString()
}
