package cart

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	Id         int
	UserId     int
	IsCheckout bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

type CartDetails struct {
	Id        int
	CartId    int
	ProductId int
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewCart(userId int, isCheckout bool, createdAt time.Time) Cart {
	return Cart{
		UserId:     userId,
		IsCheckout: isCheckout,
		CreatedAt:  createdAt,
		UpdatedAt:  createdAt,
	}
}

func newCartDetails(cartId int, productId int, quantity int, createdAt time.Time) CartDetails {
	return CartDetails{
		CartId:    cartId,
		ProductId: productId,
		Quantity:  quantity,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}
}
