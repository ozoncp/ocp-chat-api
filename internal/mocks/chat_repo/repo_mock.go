// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package chat_repo is a generated GoMock package.
package chat_repo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	chat "github.com/ozoncp/ocp-chat-api/internal/chat"
	chat_flusher "github.com/ozoncp/ocp-chat-api/internal/chat_flusher"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockRepo) Add(mess *chat.Chat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", mess)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockRepoMockRecorder) Add(mess interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRepo)(nil).Add), mess)
}

// AddBatch mocks base method.
func (m *MockRepo) AddBatch(mess []*chat.Chat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBatch", mess)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBatch indicates an expected call of AddBatch.
func (mr *MockRepoMockRecorder) AddBatch(mess interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBatch", reflect.TypeOf((*MockRepo)(nil).AddBatch), mess)
}

// DescribeByID mocks base method.
func (m *MockRepo) DescribeByID(messageID uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeByID", messageID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeByID indicates an expected call of DescribeByID.
func (mr *MockRepoMockRecorder) DescribeByID(messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeByID", reflect.TypeOf((*MockRepo)(nil).DescribeByID), messageID)
}

// GetAll mocks base method.
func (m *MockRepo) GetAll() ([]*chat.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*chat.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepoMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepo)(nil).GetAll))
}

// List mocks base method.
func (m *MockRepo) List() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRepoMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepo)(nil).List))
}

// RemoveByID mocks base method.
func (m *MockRepo) RemoveByID(messageID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveByID", messageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveByID indicates an expected call of RemoveByID.
func (mr *MockRepoMockRecorder) RemoveByID(messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveByID", reflect.TypeOf((*MockRepo)(nil).RemoveByID), messageID)
}

// MockFlusher is a mock of Flusher interface.
type MockFlusher struct {
	ctrl     *gomock.Controller
	recorder *MockFlusherMockRecorder
}

// MockFlusherMockRecorder is the mock recorder for MockFlusher.
type MockFlusherMockRecorder struct {
	mock *MockFlusher
}

// NewMockFlusher creates a new mock instance.
func NewMockFlusher(ctrl *gomock.Controller) *MockFlusher {
	mock := &MockFlusher{ctrl: ctrl}
	mock.recorder = &MockFlusherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlusher) EXPECT() *MockFlusherMockRecorder {
	return m.recorder
}

// Flush mocks base method.
func (m *MockFlusher) Flush(repo chat_flusher.FlushableChatRepo, chats []*chat.Chat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", repo, chats)
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockFlusherMockRecorder) Flush(repo, chats interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockFlusher)(nil).Flush), repo, chats)
}
