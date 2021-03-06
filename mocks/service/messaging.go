// Code generated by MockGen. DO NOT EDIT.
// Source: messaging.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	service "produce-subscription-change/service"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMessaging is a mock of Messaging interface.
type MockMessaging struct {
	ctrl     *gomock.Controller
	recorder *MockMessagingMockRecorder
}

// MockMessagingMockRecorder is the mock recorder for MockMessaging.
type MockMessagingMockRecorder struct {
	mock *MockMessaging
}

// NewMockMessaging creates a new mock instance.
func NewMockMessaging(ctrl *gomock.Controller) *MockMessaging {
	mock := &MockMessaging{ctrl: ctrl}
	mock.recorder = &MockMessagingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessaging) EXPECT() *MockMessagingMockRecorder {
	return m.recorder
}

// SendMessages mocks base method.
func (m *MockMessaging) SendMessages(messages *[]service.Message, success func(*service.Message)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessages", messages, success)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessages indicates an expected call of SendMessages.
func (mr *MockMessagingMockRecorder) SendMessages(messages, success interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessages", reflect.TypeOf((*MockMessaging)(nil).SendMessages), messages, success)
}
