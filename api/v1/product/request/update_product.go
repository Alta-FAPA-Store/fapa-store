package request

import "go-hexagonal/business/product"

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	CategoryId  int    `json:"category_id"`
}

func (req *UpdateProductRequest) ToUpsertProductSpec() *product.UpdateProductSpec {

	var uodateProductSpec product.UpdateProductSpec

	uodateProductSpec.Name = req.Name
	uodateProductSpec.Description = req.Description
	uodateProductSpec.Stock = req.Stock
	uodateProductSpec.Price = req.Price
	uodateProductSpec.CategoryId = req.CategoryId

	return &uodateProductSpec
}
