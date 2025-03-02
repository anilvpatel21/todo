package todo_test

import (
	"errors"
	"testing"
	"time"

	"example.com/todo-app/internal/todo"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// go test -v ./... to run all test cases
func TestAddAttachment(t *testing.T) {
	// Create a mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the repository
	mockRepo := todo.NewMockTodoRepository(ctrl)

	// Create the service with the mocked repository
	service := todo.NewService(mockRepo)

	// Define the test cases
	tests := []struct {
		name           string
		todoID         string
		userID         string
		fileName       string
		fileType       string
		filePath       string
		mockGetByID    func()
		mockUpdate     func()
		expectedError  error
		expectedResult *todo.Todo
	}{
		{
			name:     "Success - Add todo.Attachment",
			todoID:   "123",
			userID:   "user1",
			fileName: "testfile.jpg",
			fileType: "image/jpeg",
			filePath: "./uploads/testfile.jpg",
			mockGetByID: func() {
				mockRepo.EXPECT().GetByID("123", "user1").Return(&todo.Todo{
					ID:          "123",
					UserID:      "user1",
					Title:       "Test TODO",
					Description: "Test Description",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					Attachments: []todo.Attachment{},
				}, nil)
			},
			mockUpdate: func() {
				mockRepo.EXPECT().Update(gomock.Any()).Return(nil)
			},
			expectedError: nil,
			expectedResult: &todo.Todo{
				ID:          "123",
				UserID:      "user1",
				Title:       "Test TODO",
				Description: "Test Description",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Attachments: []todo.Attachment{
					{
						ID:       "test-attachment-id",
						FileName: "testfile.jpg",
						FileType: "image/jpeg",
						FilePath: "./uploads/testfile.jpg",
					},
				},
			},
		},
		{
			name:     "Error - todo.Todo Not Found",
			todoID:   "123",
			userID:   "user1",
			fileName: "testfile.jpg",
			fileType: "image/jpeg",
			filePath: "./uploads/testfile.jpg",
			mockGetByID: func() {
				mockRepo.EXPECT().GetByID("123", "user1").Return(nil, errors.New("todo not found"))
			},
			mockUpdate:     func() {},
			expectedError:  errors.New("todo not found"),
			expectedResult: nil,
		},
		{
			name:     "Error - Update Failed",
			todoID:   "123",
			userID:   "user1",
			fileName: "testfile.jpg",
			fileType: "image/jpeg",
			filePath: "./uploads/testfile.jpg",
			mockGetByID: func() {
				mockRepo.EXPECT().GetByID("123", "user1").Return(&todo.Todo{
					ID:          "123",
					UserID:      "user1",
					Title:       "Test TODO",
					Description: "Test Description",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					Attachments: []todo.Attachment{},
				}, nil)
			},
			mockUpdate: func() {
				mockRepo.EXPECT().Update(gomock.Any()).Return(errors.New("update failed"))
			},
			expectedError:  errors.New("update failed"),
			expectedResult: nil,
		},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mocks
			tt.mockGetByID()
			tt.mockUpdate()

			// Call the AddAttachment method
			attachment, err := service.AddAttachment(tt.todoID, tt.userID, tt.fileName, tt.fileType, tt.filePath)

			// Check the result
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult.Attachments[0].FileName, attachment.FileName)
				assert.Equal(t, tt.expectedResult.Attachments[0].FileType, attachment.FileType)
				assert.Equal(t, tt.expectedResult.Attachments[0].FilePath, attachment.FilePath)
			}
		})
	}
}
