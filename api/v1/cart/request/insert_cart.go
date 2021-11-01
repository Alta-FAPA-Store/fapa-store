package request

import "go-hexagonal/business/cart"

type InsertCartRequest struct {
	ProductId int `json:"product_id"`
}

func (req *InsertCartRequest) ToUpsertCartSpec(userId int, productId int) *cart.InsertCartSpec {
	var insertCartSpec cart.InsertCartSpec

	insertCartSpec.UserId = userId
	insertCartSpec.ProductId = productId

	return &insertCartSpec
}
