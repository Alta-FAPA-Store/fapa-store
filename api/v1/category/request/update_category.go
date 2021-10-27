package request

import "go-hexagonal/business/category"

type UpdateCategoryRequest struct {
	Name string `json:"name"`
}

func (req *UpdateCategoryRequest) ToUpsertCategorySpec() *category.UpdateCategorySpec {

	var uodateCategorySpec category.UpdateCategorySpec

	uodateCategorySpec.Name = req.Name

	return &uodateCategorySpec
}
