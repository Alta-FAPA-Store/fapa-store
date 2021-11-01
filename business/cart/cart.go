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
	Details    []CartDetailsWithProduct
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

type CartDetailsWithProduct struct {
	Id       int    `json:"cart_details_id"`
	Name     string `json:"product_name"`
	Price    int    `json:"product_price"`
	Quantity int    `json:"quantity"`
}

func NewCart(id int, userId int, isCheckout bool, createdAt time.Time) Cart {
	return Cart{
		Id:         id,
		UserId:     userId,
		IsCheckout: isCheckout,
		CreatedAt:  createdAt,
		UpdatedAt:  createdAt,
	}
}

func NewCartDetails(cartId int, productId int, quantity int, createdAt time.Time) CartDetails {
	return CartDetails{
		CartId:    cartId,
		ProductId: productId,
		Quantity:  quantity,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}
}
