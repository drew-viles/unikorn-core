// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go
//
// Generated by this command:
//
//	mockgen -source=interfaces.go -destination=mock/interfaces.go -package=mock
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	api "k8s.io/client-go/tools/clientcmd/api"
)

// MockK8SAPITester is a mock of K8SAPITester interface.
type MockK8SAPITester struct {
	ctrl     *gomock.Controller
	recorder *MockK8SAPITesterMockRecorder
}

// MockK8SAPITesterMockRecorder is the mock recorder for MockK8SAPITester.
type MockK8SAPITesterMockRecorder struct {
	mock *MockK8SAPITester
}

// NewMockK8SAPITester creates a new mock instance.
func NewMockK8SAPITester(ctrl *gomock.Controller) *MockK8SAPITester {
	mock := &MockK8SAPITester{ctrl: ctrl}
	mock.recorder = &MockK8SAPITesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockK8SAPITester) EXPECT() *MockK8SAPITesterMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockK8SAPITester) Connect(arg0 context.Context, arg1 *api.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockK8SAPITesterMockRecorder) Connect(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockK8SAPITester)(nil).Connect), arg0, arg1)
}
