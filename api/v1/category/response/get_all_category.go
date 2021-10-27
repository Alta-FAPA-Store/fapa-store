package response

import (
	"go-hexagonal/api/paginator"
	"go-hexagonal/business/category"
)

type getAllCategoryResponse struct {
	Meta      paginator.Meta        `json:"meta"`
	Categorys []GetCategoryResponse `json:"categories"`
}

//NewGetAllCategoryResponse construct GetAllCategoryResponse
func NewGetAllCategoryResponse(categories []category.Category, page int, rowPerPage int) getAllCategoryResponse {

	var (
		lenCategorys = len(categories)
	)

	getAllCategoryResponse := getAllCategoryResponse{}
	getAllCategoryResponse.Meta.BuildMeta(lenCategorys, page, rowPerPage)

	for index, value := range categories {
		if index == getAllCategoryResponse.Meta.RowPerPage {
			continue
		}

		var getCategoryResponse GetCategoryResponse

		getCategoryResponse.ID = value.CategoryID
		getCategoryResponse.Name = value.CategoryName

		getAllCategoryResponse.Categorys = append(getAllCategoryResponse.Categorys, getCategoryResponse)
	}

	if getAllCategoryResponse.Categorys == nil {
		getAllCategoryResponse.Categorys = []GetCategoryResponse{}
	}

	return getAllCategoryResponse
}
