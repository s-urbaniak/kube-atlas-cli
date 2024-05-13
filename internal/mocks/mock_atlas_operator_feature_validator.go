// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/s-urbaniak/kube-atlas-cli/internal/kubernetes/operator/features (interfaces: FeatureValidator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFeatureValidator is a mock of FeatureValidator interface.
type MockFeatureValidator struct {
	ctrl     *gomock.Controller
	recorder *MockFeatureValidatorMockRecorder
}

// MockFeatureValidatorMockRecorder is the mock recorder for MockFeatureValidator.
type MockFeatureValidatorMockRecorder struct {
	mock *MockFeatureValidator
}

// NewMockFeatureValidator creates a new mock instance.
func NewMockFeatureValidator(ctrl *gomock.Controller) *MockFeatureValidator {
	mock := &MockFeatureValidator{ctrl: ctrl}
	mock.recorder = &MockFeatureValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeatureValidator) EXPECT() *MockFeatureValidatorMockRecorder {
	return m.recorder
}

// FeatureExist mocks base method.
func (m *MockFeatureValidator) FeatureExist(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeatureExist", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// FeatureExist indicates an expected call of FeatureExist.
func (mr *MockFeatureValidatorMockRecorder) FeatureExist(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeatureExist", reflect.TypeOf((*MockFeatureValidator)(nil).FeatureExist), arg0, arg1)
}
