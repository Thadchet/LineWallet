// Code generated by MockGen. DO NOT EDIT.
// Source: transaction.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	models "line-wallet/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

// MockITransactionRepo is a mock of ITransactionRepo interface.
type MockITransactionRepo struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionRepoMockRecorder
}

// MockITransactionRepoMockRecorder is the mock recorder for MockITransactionRepo.
type MockITransactionRepoMockRecorder struct {
	mock *MockITransactionRepo
}

// NewMockITransactionRepo creates a new mock instance.
func NewMockITransactionRepo(ctrl *gomock.Controller) *MockITransactionRepo {
	mock := &MockITransactionRepo{ctrl: ctrl}
	mock.recorder = &MockITransactionRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionRepo) EXPECT() *MockITransactionRepoMockRecorder {
	return m.recorder
}

// FilterIncomeCurrentMonth mocks base method.
func (m *MockITransactionRepo) FilterIncomeCurrentMonth(line_user_id string) ([]models.Income, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterIncomeCurrentMonth", line_user_id)
	ret0, _ := ret[0].([]models.Income)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterIncomeCurrentMonth indicates an expected call of FilterIncomeCurrentMonth.
func (mr *MockITransactionRepoMockRecorder) FilterIncomeCurrentMonth(line_user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterIncomeCurrentMonth", reflect.TypeOf((*MockITransactionRepo)(nil).FilterIncomeCurrentMonth), line_user_id)
}

// FilterTransactionCurrentMonth mocks base method.
func (m *MockITransactionRepo) FilterTransactionCurrentMonth(line_user_id string) ([]models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterTransactionCurrentMonth", line_user_id)
	ret0, _ := ret[0].([]models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterTransactionCurrentMonth indicates an expected call of FilterTransactionCurrentMonth.
func (mr *MockITransactionRepoMockRecorder) FilterTransactionCurrentMonth(line_user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterTransactionCurrentMonth", reflect.TypeOf((*MockITransactionRepo)(nil).FilterTransactionCurrentMonth), line_user_id)
}

// GetIncomeByID mocks base method.
func (m *MockITransactionRepo) GetIncomeByID(ID string) (*models.Income, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncomeByID", ID)
	ret0, _ := ret[0].(*models.Income)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncomeByID indicates an expected call of GetIncomeByID.
func (mr *MockITransactionRepoMockRecorder) GetIncomeByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncomeByID", reflect.TypeOf((*MockITransactionRepo)(nil).GetIncomeByID), ID)
}

// GetTransactionByID mocks base method.
func (m *MockITransactionRepo) GetTransactionByID(ID string) (*models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByID", ID)
	ret0, _ := ret[0].(*models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByID indicates an expected call of GetTransactionByID.
func (mr *MockITransactionRepoMockRecorder) GetTransactionByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByID", reflect.TypeOf((*MockITransactionRepo)(nil).GetTransactionByID), ID)
}

// GetTransactions mocks base method.
func (m *MockITransactionRepo) GetTransactions(line_user_id string) ([]models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", line_user_id)
	ret0, _ := ret[0].([]models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockITransactionRepoMockRecorder) GetTransactions(line_user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockITransactionRepo)(nil).GetTransactions), line_user_id)
}

// Insert mocks base method.
func (m *MockITransactionRepo) Insert() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert")
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockITransactionRepoMockRecorder) Insert() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockITransactionRepo)(nil).Insert))
}

// InsertTransaction mocks base method.
func (m_2 *MockITransactionRepo) InsertTransaction(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "InsertTransaction", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertTransaction indicates an expected call of InsertTransaction.
func (mr *MockITransactionRepoMockRecorder) InsertTransaction(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTransaction", reflect.TypeOf((*MockITransactionRepo)(nil).InsertTransaction), m)
}

// UpdateIncomeByID mocks base method.
func (m *MockITransactionRepo) UpdateIncomeByID(amount, month, memo, id string) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIncomeByID", amount, month, memo, id)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIncomeByID indicates an expected call of UpdateIncomeByID.
func (mr *MockITransactionRepoMockRecorder) UpdateIncomeByID(amount, month, memo, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIncomeByID", reflect.TypeOf((*MockITransactionRepo)(nil).UpdateIncomeByID), amount, month, memo, id)
}

// UpdateTransactionByID mocks base method.
func (m *MockITransactionRepo) UpdateTransactionByID(amount, category, memo, id string) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransactionByID", amount, category, memo, id)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransactionByID indicates an expected call of UpdateTransactionByID.
func (mr *MockITransactionRepoMockRecorder) UpdateTransactionByID(amount, category, memo, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransactionByID", reflect.TypeOf((*MockITransactionRepo)(nil).UpdateTransactionByID), amount, category, memo, id)
}