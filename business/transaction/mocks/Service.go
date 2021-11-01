package mocks

import (
	"go-hexagonal/business/transaction"

	mock "github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (_m *Service) GetAllTransaction(userId int, limit int, offset int) ([]transaction.Transaction, error) {
	ret := _m.Called(userId, limit, offset)

	var r0 []transaction.Transaction
	if rf, ok := ret.Get(0).(func(int, int, int) []transaction.Transaction); ok {
		r0 = rf(userId, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transaction.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, int) error); ok {
		r1 = rf(userId, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Service) GetTransactionDetails(transactionId int) (*transaction.Transaction, error) {
	ret := _m.Called(transactionId)

	var r0 *transaction.Transaction
	if rf, ok := ret.Get(0).(func(int) *transaction.Transaction); ok {
		r0 = rf(transactionId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transaction.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(transactionId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Service) CreateTransaction(createTransactionSpec transaction.CreateTransactionSpec) (string, error) {
	ret := _m.Called(createTransactionSpec)

	var r0 *transaction.Transaction
	if rf, ok := ret.Get(0).(func(transaction.CreateTransactionSpec) *transaction.Transaction); ok {
		r0 = rf(createTransactionSpec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transaction.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(transaction.CreateTransactionSpec) error); ok {
		r1 = rf(createTransactionSpec)
	} else {
		r1 = ret.Error(0)
	}

	return r0.PaymentMethod, r1
}
