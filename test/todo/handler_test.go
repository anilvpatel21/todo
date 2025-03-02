package todo_test

import (
	"errors"
	"testing"

	"example.com/todo-app/internal/todo"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock instance of TodoService
	mockTodoService := todo.NewMockTodoService(ctrl)

	tests := []struct {
		name             string
		userID           string
		title            string
		description      string
		mockExpectations func()
		expectedError    error
		expectedResult   *todo.Todo
	}{
		{
			name:        "Create Todo Success",
			userID:      "user1",
			title:       "Test Todo",
			description: "Test Todo Description",
			mockExpectations: func() {
				mockTodoService.EXPECT().CreateTodo("user1", "Test Todo", "Test Todo Description").
					Return(&todo.Todo{ID: "1", UserID: "user1", Title: "Test Todo", Description: "Test Todo Description"}, nil)
			},
			expectedError:  nil,
			expectedResult: &todo.Todo{ID: "1", UserID: "user1", Title: "Test Todo", Description: "Test Todo Description"},
		},
		{
			name:        "Create Todo Error",
			userID:      "user1",
			title:       "Test Todo",
			description: "Test Todo Description",
			mockExpectations: func() {
				mockTodoService.EXPECT().CreateTodo("user1", "Test Todo", "Test Todo Description").
					Return(nil, errors.New("error creating todo"))
			},
			expectedError:  errors.New("error creating todo"),
			expectedResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set mock expectations
			tt.mockExpectations()

			// Use the mock service
			todo, err := mockTodoService.CreateTodo(tt.userID, tt.title, tt.description)

			// Check error and result
			if tt.expectedError != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedResult, todo)
			}
		})
	}
}
