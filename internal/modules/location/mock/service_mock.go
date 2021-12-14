// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package mock_location is a generated GoMock package.
package mock_location

import (
	domain "BrunoDM2943/via-cep-wrapper/internal/constants/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// SearchLocation mocks base method.
func (m *MockService) SearchLocation(zipCode string) (*domain.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchLocation", zipCode)
	ret0, _ := ret[0].(*domain.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchLocation indicates an expected call of SearchLocation.
func (mr *MockServiceMockRecorder) SearchLocation(zipCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchLocation", reflect.TypeOf((*MockService)(nil).SearchLocation), zipCode)
}
