// Code generated by MockGen. DO NOT EDIT.
// Source: ./flusher.go

// Package chat_repo is a generated GoMock package.
package chat_repo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	chat "github.com/ozoncp/ocp-chat-api/internal/chat"
)

// MockFlushableChatRepo is a mock of FlushableChatRepo interface.
type MockFlushableChatRepo struct {
	ctrl     *gomock.Controller
	recorder *MockFlushableChatRepoMockRecorder
}

// MockFlushableChatRepoMockRecorder is the mock recorder for MockFlushableChatRepo.
type MockFlushableChatRepoMockRecorder struct {
	mock *MockFlushableChatRepo
}

// NewMockFlushableChatRepo creates a new mock instance.
func NewMockFlushableChatRepo(ctrl *gomock.Controller) *MockFlushableChatRepo {
	mock := &MockFlushableChatRepo{ctrl: ctrl}
	mock.recorder = &MockFlushableChatRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlushableChatRepo) EXPECT() *MockFlushableChatRepoMockRecorder {
	return m.recorder
}

// AddBatch mocks base method.
func (m *MockFlushableChatRepo) AddBatch(chats []*chat.Chat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBatch", chats)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBatch indicates an expected call of AddBatch.
func (mr *MockFlushableChatRepoMockRecorder) AddBatch(chats interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBatch", reflect.TypeOf((*MockFlushableChatRepo)(nil).AddBatch), chats)
}
