package cart

import (
	"go-hexagonal/business"
	"go-hexagonal/util/validator"
	"time"
)

type InsertCartSpec struct {
	UserId    int `validate:"required"`
	ProductId int `validate:"required"`
}

type DeleteCartDetailsSpec struct {
	CartId    int `validate:"required"`
	ProductId int `validate:"required"`
}

type FindCartSpec struct {
	UserId int `validate:"required"`
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) FindCartByUserId(userId int) (*Cart, error) {
	cart, err := s.repository.FindCartByUserId(userId)

	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (s *service) InsertCart(insertCartSpec InsertCartSpec) error {
	err := validator.GetValidator().Struct(insertCartSpec)

	if err != nil {
		return business.ErrInvalidSpec
	}

	cart, err := s.repository.FindCartByUserId(insertCartSpec.UserId)

	if err != nil {
		return err
	}

	// Check cart user is empty or not. If empty then add a new one, if not empty then just add cart details
	var cartId int
	if cart == nil {
		cart := NewCart(insertCartSpec.UserId, false, time.Now())

		id, err := s.repository.InsertCart(cart)
		cartId = id

		if err != nil {
			return err
		}
	} else {
		cartId = cart.Id
	}

	// Check product in cart, if exist then just update the quantity
	isProductExist, err := s.repository.CheckCartProduct(cartId, insertCartSpec.ProductId)

	if err != nil {
		return err
	}

	if isProductExist {
		err = s.repository.UpdateCartDetailsProduct(cartId, insertCartSpec.ProductId)

		if err != nil {
			return err
		}
	} else {
		cartDetails := newCartDetails(cartId, insertCartSpec.ProductId, 1, time.Now())

		err = s.repository.InsertCartDetails(cartDetails)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) DeleteCartDetails(deleteCartDetails DeleteCartDetailsSpec) error {
	err := s.repository.DeleteCartDetails(deleteCartDetails.CartId, deleteCartDetails.ProductId)

	if err != nil {
		return err
	}

	return nil
}
