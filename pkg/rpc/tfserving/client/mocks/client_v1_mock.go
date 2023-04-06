// Code generated by MockGen. DO NOT EDIT.
// Source: client_v1.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfserving "d7y.io/api/pkg/apis/tfserving/v1"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockV1 is a mock of V1 interface.
type MockV1 struct {
	ctrl     *gomock.Controller
	recorder *MockV1MockRecorder
}

// MockV1MockRecorder is the mock recorder for MockV1.
type MockV1MockRecorder struct {
	mock *MockV1
}

// NewMockV1 creates a new mock instance.
func NewMockV1(ctrl *gomock.Controller) *MockV1 {
	mock := &MockV1{ctrl: ctrl}
	mock.recorder = &MockV1MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockV1) EXPECT() *MockV1MockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockV1) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockV1MockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockV1)(nil).Close))
}

// Predict mocks base method.
func (m *MockV1) Predict(arg0 context.Context, arg1 *tfserving.PredictRequest, arg2 ...grpc.CallOption) (*tfserving.PredictResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Predict", varargs...)
	ret0, _ := ret[0].(*tfserving.PredictResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Predict indicates an expected call of Predict.
func (mr *MockV1MockRecorder) Predict(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Predict", reflect.TypeOf((*MockV1)(nil).Predict), varargs...)
}