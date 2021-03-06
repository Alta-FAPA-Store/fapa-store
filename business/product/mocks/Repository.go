package mocks

import (
	product "go-hexagonal/business/product"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

func (_m *Repository) FindProductByID(id int) (*product.Product, error) {
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

func (_m *Repository) FindAllProduct(skip, rowPerPage int, categoryParam, nameParam string) ([]product.Product, error) {
	ret := _m.Called(skip, rowPerPage, categoryParam, nameParam)

	var r0 *[]product.Product
	if rf, ok := ret.Get(0).(func(int, int, string, string) *[]product.Product); ok {
		r0 = rf(skip, rowPerPage, categoryParam, nameParam)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string, string) error); ok {
		r1 = rf(skip, rowPerPage, categoryParam, nameParam)
	} else {
		r1 = ret.Error(1)
	}

	return *r0, r1
}

func (_m *Repository) InsertProduct(_a0 product.Product) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(product.Product) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Repository) UpdateProduct(_a0 product.Product) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(product.Product) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Repository) DeleteProduct(productId int) error {
	ret := _m.Called(productId)

	var r0 error
	if rf, ok := ret.Get(0).(func(productId int) error); ok {
		r0 = rf(productId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
