package request

import "go-hexagonal/business/cart"

type UpdateCartDetailsResponse struct {
	CartId    int `json:"cart_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func (req *UpdateCartDetailsResponse) ToUpsertUpdateCartDetailsSpec() *cart.UpdateCartDetailsSpec {
	var updateCartDetailsSpec cart.UpdateCartDetailsSpec

	updateCartDetailsSpec.CartId = req.CartId
	updateCartDetailsSpec.ProductId = req.ProductId
	updateCartDetailsSpec.Quantity = req.Quantity

	return &updateCartDetailsSpec
}
