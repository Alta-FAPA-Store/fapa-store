package response

import (
	"go-hexagonal/business/category"
)

//GetPetResponse Get pet by ID response payload
type GetCategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//NewGetPetResponse construct GetPetResponse
func NewGetCategoryResponse(category category.Category) *GetCategoryResponse {
	var getCategoryResponse GetCategoryResponse

	getCategoryResponse.ID = category.CategoryID
	getCategoryResponse.Name = category.CategoryName

	return &getCategoryResponse
}
