package response

import (
	"go-hexagonal/business/cart"
)

type GetCartResponse struct {
	Id         int                      `json:"id"`
	UserId     int                      `json:"user_id"`
	IsCheckout bool                     `json:"is_checkout"`
	Details    []GetCartDetailsResponse `json:"details"`
}

type GetCartDetailsResponse struct {
	Id        int `json:"id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func NewGetCartResponse(cart cart.Cart) *GetCartResponse {
	var cartResponse GetCartResponse

	cartResponse.Id = cart.Id
	cartResponse.UserId = cart.UserId
	cartResponse.IsCheckout = cart.IsCheckout

	return &cartResponse
}
