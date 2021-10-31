package mocks

import (
	product "go-hexagonal/business/product"

	mock "github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (_m *Service) FindProductByID(id int) (*product.Product, error) {
	ret := _m.Called(id)

	var r0 *product.Product
	if rf, ok := ret.Get(0).(func(int) *product.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
