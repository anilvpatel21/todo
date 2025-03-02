// Package todo is a generated GoMock package.
package todo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTodoRepository is a mock of TodoRepository interface.
type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

// MockTodoRepositoryMockRecorder is the mock recorder for MockTodoRepository.
type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

// NewMockTodoRepository creates a new mock instance.
func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTodoRepository) Create(todo Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", todo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTodoRepositoryMockRecorder) Create(todo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoRepository)(nil).Create), todo)
}

// Delete mocks base method.
func (m *MockTodoRepository) Delete(id, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoRepositoryMockRecorder) Delete(id, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoRepository)(nil).Delete), id, userID)
}

// GetByID mocks base method.
func (m *MockTodoRepository) GetByID(id, userID string) (*Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id, userID)
	ret0, _ := ret[0].(*Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockTodoRepositoryMockRecorder) GetByID(id, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTodoRepository)(nil).GetByID), id, userID)
}

// GetAllByUserID mocks base method.
func (m *MockTodoRepository) GetAllByUserID(userID string) ([]Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByUserID", userID)
	ret0, _ := ret[0].([]Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByUserID indicates an expected call of GetAllByUserID.
func (mr *MockTodoRepositoryMockRecorder) GetAllByUserID(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByUserID", reflect.TypeOf((*MockTodoRepository)(nil).GetAllByUserID), userID)
}

// Update mocks base method.
func (m *MockTodoRepository) Update(todo Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", todo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTodoRepositoryMockRecorder) Update(todo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoRepository)(nil).Update), todo)
}
