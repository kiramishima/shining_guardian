// Code generated by MockGen. DO NOT EDIT.
// Source: .\src\ports\repository\auth_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/kiramishima/shining_guardian/src/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockIAuthRepository is a mock of IAuthRepository interface.
type MockIAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthRepositoryMockRecorder
}

// MockIAuthRepositoryMockRecorder is the mock recorder for MockIAuthRepository.
type MockIAuthRepositoryMockRecorder struct {
	mock *MockIAuthRepository
}

// NewMockIAuthRepository creates a new mock instance.
func NewMockIAuthRepository(ctrl *gomock.Controller) *MockIAuthRepository {
	mock := &MockIAuthRepository{ctrl: ctrl}
	mock.recorder = &MockIAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthRepository) EXPECT() *MockIAuthRepositoryMockRecorder {
	return m.recorder
}

// FindByCredentials mocks base method.
func (m *MockIAuthRepository) FindByCredentials(ctx context.Context, data *domain.AuthRequest) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCredentials", ctx, data)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCredentials indicates an expected call of FindByCredentials.
func (mr *MockIAuthRepositoryMockRecorder) FindByCredentials(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCredentials", reflect.TypeOf((*MockIAuthRepository)(nil).FindByCredentials), ctx, data)
}

// Register mocks base method.
func (m *MockIAuthRepository) Register(ctx context.Context, registerReq *domain.RegisterRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, registerReq)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockIAuthRepositoryMockRecorder) Register(ctx, registerReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIAuthRepository)(nil).Register), ctx, registerReq)
}
