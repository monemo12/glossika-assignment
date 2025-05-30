// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/user_repo.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	model "glossika-assignment/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// CheckUserExists mocks base method.
func (m *MockIUserRepository) CheckUserExists(ctx context.Context, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserExists", ctx, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserExists indicates an expected call of CheckUserExists.
func (mr *MockIUserRepositoryMockRecorder) CheckUserExists(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserExists", reflect.TypeOf((*MockIUserRepository)(nil).CheckUserExists), ctx, email)
}

// CreateUser mocks base method.
func (m *MockIUserRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserRepositoryMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserRepository)(nil).CreateUser), ctx, user)
}

// GetUserByEmail mocks base method.
func (m *MockIUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockIUserRepositoryMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockIUserRepository)(nil).GetUserByEmail), ctx, email)
}

// GetUserByVerificationToken mocks base method.
func (m *MockIUserRepository) GetUserByVerificationToken(ctx context.Context, token string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByVerificationToken", ctx, token)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByVerificationToken indicates an expected call of GetUserByVerificationToken.
func (mr *MockIUserRepositoryMockRecorder) GetUserByVerificationToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByVerificationToken", reflect.TypeOf((*MockIUserRepository)(nil).GetUserByVerificationToken), ctx, token)
}

// UpdateUserVerification mocks base method.
func (m *MockIUserRepository) UpdateUserVerification(ctx context.Context, userID string, verified bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserVerification", ctx, userID, verified)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserVerification indicates an expected call of UpdateUserVerification.
func (mr *MockIUserRepositoryMockRecorder) UpdateUserVerification(ctx, userID, verified interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserVerification", reflect.TypeOf((*MockIUserRepository)(nil).UpdateUserVerification), ctx, userID, verified)
}
