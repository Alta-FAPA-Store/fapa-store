package request

import (
	"go-hexagonal/business/category"
)

type InsertCategoryRequest struct {
	Name string `json:"name"`
}

func (req *InsertCategoryRequest) ToUpsertCategorySpec() *category.InsertCategorySpec {

	var insertCategorySpec category.InsertCategorySpec

	insertCategorySpec.Name = req.Name

	return &insertCategorySpec
}
