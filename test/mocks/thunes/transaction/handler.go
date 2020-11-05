// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/thunes/transaction/handler.go

// Package mock_transaction is a generated GoMock package.
package mock_transaction

import (
	transaction "fghpdf.me/thunes_homework/internal/pkg/thunes/transaction"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockServer is a mock of Server interface
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockServer) Create(quotationId uint64, params *transaction.CreateParams) (*transaction.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", quotationId, params)
	ret0, _ := ret[0].(*transaction.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockServerMockRecorder) Create(quotationId, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServer)(nil).Create), quotationId, params)
}

// Confirm mocks base method
func (m *MockServer) Confirm(transactionId uint64) (*transaction.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Confirm", transactionId)
	ret0, _ := ret[0].(*transaction.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Confirm indicates an expected call of Confirm
func (mr *MockServerMockRecorder) Confirm(transactionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Confirm", reflect.TypeOf((*MockServer)(nil).Confirm), transactionId)
}
