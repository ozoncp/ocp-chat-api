// Code generated by MockGen. DO NOT EDIT.
// Source: ./chat_api.go

// Package chat_api is a generated GoMock package.
package chat_api

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	chat "github.com/ozoncp/ocp-chat-api/internal/chat"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateChat mocks base method.
func (m *MockService) CreateChat(ctx context.Context, classroom uint64, link string) (*chat.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChat", ctx, classroom, link)
	ret0, _ := ret[0].(*chat.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChat indicates an expected call of CreateChat.
func (mr *MockServiceMockRecorder) CreateChat(ctx, classroom, link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChat", reflect.TypeOf((*MockService)(nil).CreateChat), ctx, classroom, link)
}

// CreateMultipleChat mocks base method.
func (m *MockService) CreateMultipleChat(ctx context.Context, classroom []uint64, link []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMultipleChat", ctx, classroom, link)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMultipleChat indicates an expected call of CreateMultipleChat.
func (mr *MockServiceMockRecorder) CreateMultipleChat(ctx, classroom, link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMultipleChat", reflect.TypeOf((*MockService)(nil).CreateMultipleChat), ctx, classroom, link)
}

// DescribeChat mocks base method.
func (m *MockService) DescribeChat(ctx context.Context, id uint64) (*chat.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeChat", ctx, id)
	ret0, _ := ret[0].(*chat.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeChat indicates an expected call of DescribeChat.
func (mr *MockServiceMockRecorder) DescribeChat(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeChat", reflect.TypeOf((*MockService)(nil).DescribeChat), ctx, id)
}

// ListChats mocks base method.
func (m *MockService) ListChats(ctx context.Context) ([]*chat.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChats", ctx)
	ret0, _ := ret[0].([]*chat.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListChats indicates an expected call of ListChats.
func (mr *MockServiceMockRecorder) ListChats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChats", reflect.TypeOf((*MockService)(nil).ListChats), ctx)
}

// RemoveChat mocks base method.
func (m *MockService) RemoveChat(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveChat", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveChat indicates an expected call of RemoveChat.
func (mr *MockServiceMockRecorder) RemoveChat(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveChat", reflect.TypeOf((*MockService)(nil).RemoveChat), ctx, id)
}
