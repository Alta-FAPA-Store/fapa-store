package mocks

import (
	cart "go-hexagonal/business/cart"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

func (_m *Service) FindCartByUserId(userId int) (*cart.Cart, error) {
	ret := _m.Called(userId)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(int) *cart.Cart); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Service) InsertCart(insertCartSpec cart.InsertCartSpec) (int, error) {
	ret := _m.Called(insertCartSpec)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(cart.InsertCartSpec) *cart.Cart); ok {
		r0 = rf(insertCartSpec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(cart.InsertCartSpec) error); ok {
		r1 = rf(insertCartSpec)
	} else {
		r1 = ret.Error(0)
	}

	return r0.Id, r1
}

func (_m *Service) DeleteCartDetails(cartId int, productId int) error {
	ret := _m.Called(cartId, productId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(cartId, productId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Service) UpdateQuantityCartDetails(cartId int, productId int, quantity int) error {
	ret := _m.Called(cartId, productId, quantity)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, int) error); ok {
		r0 = rf(cartId, productId, quantity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}