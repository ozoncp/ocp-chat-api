// Code generated by MockGen. DO NOT EDIT.
// Source: ./flusher.go

// Package chat_repo is a generated GoMock package.
package chat_repo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	chat "github.com/ozoncp/ocp-chat-api/internal/chat"
)

// MockChatRepo is a mock of ChatRepo interface.
type MockChatRepo struct {
	ctrl     *gomock.Controller
	recorder *MockChatRepoMockRecorder
}

// MockChatRepoMockRecorder is the mock recorder for MockChatRepo.
type MockChatRepoMockRecorder struct {
	mock *MockChatRepo
}

// NewMockChatRepo creates a new mock instance.
func NewMockChatRepo(ctrl *gomock.Controller) *MockChatRepo {
	mock := &MockChatRepo{ctrl: ctrl}
	mock.recorder = &MockChatRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatRepo) EXPECT() *MockChatRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockChatRepo) Add(mess *chat.Chat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", mess)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockChatRepoMockRecorder) Add(mess interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockChatRepo)(nil).Add), mess)
}

// AddBatch mocks base method.
func (m *MockChatRepo) AddBatch(mess []*chat.Chat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBatch", mess)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBatch indicates an expected call of AddBatch.
func (mr *MockChatRepoMockRecorder) AddBatch(mess interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBatch", reflect.TypeOf((*MockChatRepo)(nil).AddBatch), mess)
}

// DescribeByID mocks base method.
func (m *MockChatRepo) DescribeByID(messageID uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeByID", messageID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeByID indicates an expected call of DescribeByID.
func (mr *MockChatRepoMockRecorder) DescribeByID(messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeByID", reflect.TypeOf((*MockChatRepo)(nil).DescribeByID), messageID)
}

// GetAll mocks base method.
func (m *MockChatRepo) GetAll() ([]*chat.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*chat.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockChatRepoMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockChatRepo)(nil).GetAll))
}

// List mocks base method.
func (m *MockChatRepo) List() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockChatRepoMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockChatRepo)(nil).List))
}

// RemoveByID mocks base method.
func (m *MockChatRepo) RemoveByID(messageID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveByID", messageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveByID indicates an expected call of RemoveByID.
func (mr *MockChatRepoMockRecorder) RemoveByID(messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveByID", reflect.TypeOf((*MockChatRepo)(nil).RemoveByID), messageID)
}
