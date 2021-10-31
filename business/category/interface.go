package category

type Service interface {
	//FindCategoryByID If data not found will return nil without error
	FindCategoryByID(id int) (*Category, error)

	FindAllCategory(skip int, rowPerPage int) ([]Category, error)

	//InsertCategory Insert new Category into storage
	InsertCategory(insertCategorySpec InsertCategorySpec, createdBy string) error

	UpdateCategory(id int, updateCategorySpec UpdateCategorySpec) error

	DeleteCategory(categoryId int) error
}

type Repository interface {
	//FindCategoryByID If data not found will return nil without error
	FindCategoryByID(id int) (*Category, error)

	FindAllCategory(skip int, rowPerPage int) ([]Category, error)

	InsertCategory(category Category) error

	UpdateCategory(category Category) error

	DeleteCategory(categoryId int) error
}
