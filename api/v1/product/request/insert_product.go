package request

import (
	"go-hexagonal/business/product"
)

type InsertProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	CategoryID  int    `json:"category_id"`
}

func (req *InsertProductRequest) ToUpsertProductSpec() *product.InsertProductSpec {

	var insertProductSpec product.InsertProductSpec

	insertProductSpec.Name = req.Name
	insertProductSpec.Description = req.Description
	insertProductSpec.Stock = req.Stock
	insertProductSpec.Price = req.Price
	insertProductSpec.CategoryId = req.CategoryID

	return &insertProductSpec
}
