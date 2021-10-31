package response

import (
	"go-hexagonal/api/paginator"
	"go-hexagonal/business/product"
)

type getAllProductResponse struct {
	Meta     paginator.Meta       `json:"meta"`
	Products []GetProductResponse `json:"products"`
}

//NewGetAllProductResponse construct GetAllProductResponse
func NewGetAllProductResponse(products []product.Product, page int, rowPerPage int) getAllProductResponse {

	var (
		lenProducts = len(products)
	)

	getAllProductResponse := getAllProductResponse{}
	getAllProductResponse.Meta.BuildMeta(lenProducts, page, rowPerPage)

	for index, value := range products {
		if index == getAllProductResponse.Meta.RowPerPage {
			continue
		}

		var getProductResponse GetProductResponse

		getProductResponse.ID = value.ID
		getProductResponse.CategoryID = value.CategoryID
		getProductResponse.Name = value.Name
		getProductResponse.Price = value.Price
		getProductResponse.Stock = value.Stock
		getProductResponse.Slug = value.Slug
		getProductResponse.Description = value.Description

		getAllProductResponse.Products = append(getAllProductResponse.Products, getProductResponse)
	}

	if getAllProductResponse.Products == nil {
		getAllProductResponse.Products = []GetProductResponse{}
	}

	return getAllProductResponse
}
