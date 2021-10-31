package request

import "go-hexagonal/business/cart"

type DeleteCartDetailsRequest struct {
	CartId    int `json:"cart_id"`
	ProductId int `json:"product_id"`
}

func (req *DeleteCartDetailsRequest) ToUpsetDeleteCartDetailsSpec(cartId int, productId int) *cart.DeleteCartDetailsSpec {
	var deleteCartDetailsSpec cart.DeleteCartDetailsSpec

	deleteCartDetailsSpec.CartId = cartId
	deleteCartDetailsSpec.ProductId = productId

	return &deleteCartDetailsSpec
}
