// Code generated by MockGen. DO NOT EDIT.
// Source: squaremicroservices/app (interfaces: Square)

// Package mocksquareapp is a generated GoMock package.
package mocksquareapp

import (
	reflect "reflect"
	app "squaremicroservices/app"

	gomock "github.com/golang/mock/gomock"
	resources "github.com/longvu727/FootballSquaresLibs/util/resources"
)

// MockSquare is a mock of Square interface.
type MockSquare struct {
	ctrl     *gomock.Controller
	recorder *MockSquareMockRecorder
}

// MockSquareMockRecorder is the mock recorder for MockSquare.
type MockSquareMockRecorder struct {
	mock *MockSquare
}

// NewMockSquare creates a new mock instance.
func NewMockSquare(ctrl *gomock.Controller) *MockSquare {
	mock := &MockSquare{ctrl: ctrl}
	mock.recorder = &MockSquareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSquare) EXPECT() *MockSquareMockRecorder {
	return m.recorder
}

// CreateDBSquare mocks base method.
func (m *MockSquare) CreateDBSquare(arg0 app.CreateSquareParams, arg1 *resources.Resources) (*app.CreateSquareResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDBSquare", arg0, arg1)
	ret0, _ := ret[0].(*app.CreateSquareResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDBSquare indicates an expected call of CreateDBSquare.
func (mr *MockSquareMockRecorder) CreateDBSquare(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDBSquare", reflect.TypeOf((*MockSquare)(nil).CreateDBSquare), arg0, arg1)
}

// GetDBSquare mocks base method.
func (m *MockSquare) GetDBSquare(arg0 app.GetSquareParams, arg1 *resources.Resources) (*app.GetSquareResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBSquare", arg0, arg1)
	ret0, _ := ret[0].(*app.GetSquareResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDBSquare indicates an expected call of GetDBSquare.
func (mr *MockSquareMockRecorder) GetDBSquare(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBSquare", reflect.TypeOf((*MockSquare)(nil).GetDBSquare), arg0, arg1)
}
