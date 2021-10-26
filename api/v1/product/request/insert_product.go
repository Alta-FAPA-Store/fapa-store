package request

import (
	"go-hexagonal/business/product"
)

type InsertProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	CategoryID  int    `json:"category"`
}

//ToUpsertPetSpec convert into Pet.UpsertPetSpec object
func (req *InsertProductRequest) ToUpsertProductSpec(userID int) *product.InsertProductSpec {

	var insertProductSpec product.InsertProductSpec

	insertProductSpec.Name = req.Name
	insertProductSpec.Description = req.Description
	insertProductSpec.Stock = req.Stock
	insertProductSpec.Price = req.Price
	insertProductSpec.Category = req.CategoryID

	return &insertProductSpec
}
