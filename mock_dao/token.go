// Code generated by MockGen. DO NOT EDIT.
// Source: token.go

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dao "github.com/joschahenningsen/TUM-Live/dao"
	model "github.com/joschahenningsen/TUM-Live/model"
)

// MockTokenDao is a mock of TokenDao interface.
type MockTokenDao struct {
	ctrl     *gomock.Controller
	recorder *MockTokenDaoMockRecorder
}

// MockTokenDaoMockRecorder is the mock recorder for MockTokenDao.
type MockTokenDaoMockRecorder struct {
	mock *MockTokenDao
}

// NewMockTokenDao creates a new mock instance.
func NewMockTokenDao(ctrl *gomock.Controller) *MockTokenDao {
	mock := &MockTokenDao{ctrl: ctrl}
	mock.recorder = &MockTokenDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenDao) EXPECT() *MockTokenDaoMockRecorder {
	return m.recorder
}

// AddToken mocks base method.
func (m *MockTokenDao) AddToken(ctx context.Context, token model.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToken", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToken indicates an expected call of AddToken.
func (mr *MockTokenDaoMockRecorder) AddToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToken", reflect.TypeOf((*MockTokenDao)(nil).AddToken), ctx, token)
}

// DeleteToken mocks base method.
func (m *MockTokenDao) DeleteToken(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteToken", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteToken indicates an expected call of DeleteToken.
func (mr *MockTokenDaoMockRecorder) DeleteToken(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteToken", reflect.TypeOf((*MockTokenDao)(nil).DeleteToken), ctx, id)
}

// GetAllTokens mocks base method.
func (m *MockTokenDao) GetAllTokens(ctx context.Context) ([]dao.AllTokensDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTokens", ctx)
	ret0, _ := ret[0].([]dao.AllTokensDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTokens indicates an expected call of GetAllTokens.
func (mr *MockTokenDaoMockRecorder) GetAllTokens(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTokens", reflect.TypeOf((*MockTokenDao)(nil).GetAllTokens), ctx)
}

// GetToken mocks base method.
func (m *MockTokenDao) GetToken(ctx context.Context, token string) (model.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", ctx, token)
	ret0, _ := ret[0].(model.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockTokenDaoMockRecorder) GetToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockTokenDao)(nil).GetToken), ctx, token)
}

// TokenUsed mocks base method.
func (m *MockTokenDao) TokenUsed(ctx context.Context, token model.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenUsed", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// TokenUsed indicates an expected call of TokenUsed.
func (mr *MockTokenDaoMockRecorder) TokenUsed(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenUsed", reflect.TypeOf((*MockTokenDao)(nil).TokenUsed), ctx, token)
}
