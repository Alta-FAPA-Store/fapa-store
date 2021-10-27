package product

type Service interface {
	//FindProductByID If data not found will return nil without error
	FindProductByID(id int) (*Product, error)

	FindAllProduct(skip int, rowPerPage int, categoryParam, nameParam string) ([]Product, error)

	//InsertProduct Insert new Product into storage
	InsertProduct(insertProductSpec InsertProductSpec, createdBy string) error

	UpdateProduct(id int, updateProductSpec UpdateProductSpec) error

	DeleteProduct(productId int) error
}

type Repository interface {
	//FindProductByID If data not found will return nil without error
	FindProductByID(id int) (*Product, error)

	FindAllProduct(skip int, rowPerPage int, categoryParam, nameParam string) ([]Product, error)

	InsertProduct(product Product) error

	UpdateProduct(product Product) error

	DeleteProduct(productId int) error
}
