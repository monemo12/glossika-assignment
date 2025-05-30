// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/recommendation_repo.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	model "glossika-assignment/internal/model"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockIRecommendationRepository is a mock of IRecommendationRepository interface.
type MockIRecommendationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRecommendationRepositoryMockRecorder
}

// MockIRecommendationRepositoryMockRecorder is the mock recorder for MockIRecommendationRepository.
type MockIRecommendationRepositoryMockRecorder struct {
	mock *MockIRecommendationRepository
}

// NewMockIRecommendationRepository creates a new mock instance.
func NewMockIRecommendationRepository(ctrl *gomock.Controller) *MockIRecommendationRepository {
	mock := &MockIRecommendationRepository{ctrl: ctrl}
	mock.recorder = &MockIRecommendationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRecommendationRepository) EXPECT() *MockIRecommendationRepositoryMockRecorder {
	return m.recorder
}

// FetchItemsByPagination mocks base method.
func (m *MockIRecommendationRepository) FetchItemsByPagination(ctx context.Context, limit, offset int) ([]*model.Recommendation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchItemsByPagination", ctx, limit, offset)
	ret0, _ := ret[0].([]*model.Recommendation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchItemsByPagination indicates an expected call of FetchItemsByPagination.
func (mr *MockIRecommendationRepositoryMockRecorder) FetchItemsByPagination(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchItemsByPagination", reflect.TypeOf((*MockIRecommendationRepository)(nil).FetchItemsByPagination), ctx, limit, offset)
}

// FetchItemsCount mocks base method.
func (m *MockIRecommendationRepository) FetchItemsCount(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchItemsCount", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchItemsCount indicates an expected call of FetchItemsCount.
func (mr *MockIRecommendationRepositoryMockRecorder) FetchItemsCount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchItemsCount", reflect.TypeOf((*MockIRecommendationRepository)(nil).FetchItemsCount), ctx)
}

// cacheItemsToRedis mocks base method.
func (m *MockIRecommendationRepository) cacheItemsToRedis(ctx context.Context, items []*model.Recommendation, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "cacheItemsToRedis", ctx, items, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// cacheItemsToRedis indicates an expected call of cacheItemsToRedis.
func (mr *MockIRecommendationRepositoryMockRecorder) cacheItemsToRedis(ctx, items, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cacheItemsToRedis", reflect.TypeOf((*MockIRecommendationRepository)(nil).cacheItemsToRedis), ctx, items, ttl)
}

// fetchItemsCountFromDB mocks base method.
func (m *MockIRecommendationRepository) fetchItemsCountFromDB(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fetchItemsCountFromDB", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fetchItemsCountFromDB indicates an expected call of fetchItemsCountFromDB.
func (mr *MockIRecommendationRepositoryMockRecorder) fetchItemsCountFromDB(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fetchItemsCountFromDB", reflect.TypeOf((*MockIRecommendationRepository)(nil).fetchItemsCountFromDB), ctx)
}

// fetchItemsCountFromRedis mocks base method.
func (m *MockIRecommendationRepository) fetchItemsCountFromRedis(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fetchItemsCountFromRedis", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fetchItemsCountFromRedis indicates an expected call of fetchItemsCountFromRedis.
func (mr *MockIRecommendationRepositoryMockRecorder) fetchItemsCountFromRedis(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fetchItemsCountFromRedis", reflect.TypeOf((*MockIRecommendationRepository)(nil).fetchItemsCountFromRedis), ctx)
}

// fetchItemsFromDB mocks base method.
func (m *MockIRecommendationRepository) fetchItemsFromDB(ctx context.Context, limit, offset int) ([]*model.Recommendation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fetchItemsFromDB", ctx, limit, offset)
	ret0, _ := ret[0].([]*model.Recommendation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fetchItemsFromDB indicates an expected call of fetchItemsFromDB.
func (mr *MockIRecommendationRepositoryMockRecorder) fetchItemsFromDB(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fetchItemsFromDB", reflect.TypeOf((*MockIRecommendationRepository)(nil).fetchItemsFromDB), ctx, limit, offset)
}

// fetchItemsFromRedis mocks base method.
func (m *MockIRecommendationRepository) fetchItemsFromRedis(ctx context.Context, limit, offset int) ([]*model.Recommendation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fetchItemsFromRedis", ctx, limit, offset)
	ret0, _ := ret[0].([]*model.Recommendation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fetchItemsFromRedis indicates an expected call of fetchItemsFromRedis.
func (mr *MockIRecommendationRepositoryMockRecorder) fetchItemsFromRedis(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fetchItemsFromRedis", reflect.TypeOf((*MockIRecommendationRepository)(nil).fetchItemsFromRedis), ctx, limit, offset)
}
