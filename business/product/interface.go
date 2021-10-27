package product

type Service interface {
	//FindProductByID If data not found will return nil without error
	FindProductByID(id int) (*Product, error)

	//InsertProduct Insert new Product into storage
	InsertProduct(insertProductSpec InsertProductSpec, createdBy string) error

	//UpdateProduct(id int, updateProductSpec UpdateProductSpec, modifiedBy string) error
}

type Repository interface {
	//FindProductByID If data not found will return nil without error
	FindProductByID(id int) (*Product, error)

	InsertProduct(product Product) error

	//UpdateProduct(id int, updateProductSpec UpdateProductSpec, modifiedBy string) error
}
