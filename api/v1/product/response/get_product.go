package response

import (
	"go-hexagonal/business/product"
)

//GetPetResponse Get pet by ID response payload
type GetProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Slug        string `json:"slug"`
	CategoryID  int    `json:"category_id"`
}

//NewGetPetResponse construct GetPetResponse
func NewGetProductResponse(product product.Product) *GetProductResponse {
	var getProductResponse GetProductResponse

	getProductResponse.ID = product.ID
	getProductResponse.Name = product.Name
	getProductResponse.Description = product.Description
	getProductResponse.Stock = product.Stock
	getProductResponse.Price = product.Price
	getProductResponse.Slug = product.Slug
	getProductResponse.CategoryID = product.CategoryID

	return &getProductResponse
}
