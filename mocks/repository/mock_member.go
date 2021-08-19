// Code generated by MockGen. DO NOT EDIT.
// Source: member.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	models "line-wallet/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIMemberRepo is a mock of IMemberRepo interface.
type MockIMemberRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIMemberRepoMockRecorder
}

// MockIMemberRepoMockRecorder is the mock recorder for MockIMemberRepo.
type MockIMemberRepoMockRecorder struct {
	mock *MockIMemberRepo
}

// NewMockIMemberRepo creates a new mock instance.
func NewMockIMemberRepo(ctrl *gomock.Controller) *MockIMemberRepo {
	mock := &MockIMemberRepo{ctrl: ctrl}
	mock.recorder = &MockIMemberRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMemberRepo) EXPECT() *MockIMemberRepoMockRecorder {
	return m.recorder
}

// CreateMember mocks base method.
func (m *MockIMemberRepo) CreateMember(member models.Member) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMember", member)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMember indicates an expected call of CreateMember.
func (mr *MockIMemberRepoMockRecorder) CreateMember(member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMember", reflect.TypeOf((*MockIMemberRepo)(nil).CreateMember), member)
}

// FindMemberByLineUserID mocks base method.
func (m *MockIMemberRepo) FindMemberByLineUserID(line_use_id string) (*models.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMemberByLineUserID", line_use_id)
	ret0, _ := ret[0].(*models.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMemberByLineUserID indicates an expected call of FindMemberByLineUserID.
func (mr *MockIMemberRepoMockRecorder) FindMemberByLineUserID(line_use_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMemberByLineUserID", reflect.TypeOf((*MockIMemberRepo)(nil).FindMemberByLineUserID), line_use_id)
}

// UpdateRemainingBalance mocks base method.
func (m *MockIMemberRepo) UpdateRemainingBalance(line_use_id string, amount float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRemainingBalance", line_use_id, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRemainingBalance indicates an expected call of UpdateRemainingBalance.
func (mr *MockIMemberRepoMockRecorder) UpdateRemainingBalance(line_use_id, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRemainingBalance", reflect.TypeOf((*MockIMemberRepo)(nil).UpdateRemainingBalance), line_use_id, amount)
}
