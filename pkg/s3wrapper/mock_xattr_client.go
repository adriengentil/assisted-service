// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/pkg/s3wrapper (interfaces: XattrClient)

// Package s3wrapper is a generated GoMock package.
package s3wrapper

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockXattrClient is a mock of XattrClient interface.
type MockXattrClient struct {
	ctrl     *gomock.Controller
	recorder *MockXattrClientMockRecorder
}

// MockXattrClientMockRecorder is the mock recorder for MockXattrClient.
type MockXattrClientMockRecorder struct {
	mock *MockXattrClient
}

// NewMockXattrClient creates a new mock instance.
func NewMockXattrClient(ctrl *gomock.Controller) *MockXattrClient {
	mock := &MockXattrClient{ctrl: ctrl}
	mock.recorder = &MockXattrClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockXattrClient) EXPECT() *MockXattrClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockXattrClient) Get(arg0, arg1 string) (string, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockXattrClientMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockXattrClient)(nil).Get), arg0, arg1)
}

// IsSupported mocks base method.
func (m *MockXattrClient) IsSupported() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSupported")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSupported indicates an expected call of IsSupported.
func (mr *MockXattrClientMockRecorder) IsSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSupported", reflect.TypeOf((*MockXattrClient)(nil).IsSupported))
}

// List mocks base method.
func (m *MockXattrClient) List(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockXattrClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockXattrClient)(nil).List), arg0)
}

// RemoveAll mocks base method.
func (m *MockXattrClient) RemoveAll(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll.
func (mr *MockXattrClientMockRecorder) RemoveAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockXattrClient)(nil).RemoveAll), arg0)
}

// Set mocks base method.
func (m *MockXattrClient) Set(arg0, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockXattrClientMockRecorder) Set(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockXattrClient)(nil).Set), arg0, arg1, arg2, arg3)
}
