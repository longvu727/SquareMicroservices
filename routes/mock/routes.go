// Code generated by MockGen. DO NOT EDIT.
// Source: squaremicroservices/routes (interfaces: RoutesInterface)

// Package mockroutes is a generated GoMock package.
package mockroutes

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	resources "github.com/longvu727/FootballSquaresLibs/util/resources"
)

// MockRoutesInterface is a mock of RoutesInterface interface.
type MockRoutesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRoutesInterfaceMockRecorder
}

// MockRoutesInterfaceMockRecorder is the mock recorder for MockRoutesInterface.
type MockRoutesInterfaceMockRecorder struct {
	mock *MockRoutesInterface
}

// NewMockRoutesInterface creates a new mock instance.
func NewMockRoutesInterface(ctrl *gomock.Controller) *MockRoutesInterface {
	mock := &MockRoutesInterface{ctrl: ctrl}
	mock.recorder = &MockRoutesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoutesInterface) EXPECT() *MockRoutesInterfaceMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockRoutesInterface) Register(arg0 *resources.Resources) *http.ServeMux {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(*http.ServeMux)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockRoutesInterfaceMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRoutesInterface)(nil).Register), arg0)
}
