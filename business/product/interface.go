package product

type Service interface {
	//FindProductByID If data not found will return nil without error
	FindProductByID(id int, userID int) (*Product, error)

	//InsertProduct Insert new Pet into storage
	InsertProduct(insertProductSpec InsertProductSpec, createdBy string) error
}

type Repository interface {
	//FindProductByID If data not found will return nil without error
	FindProductByID(id int, userID int) (*Product, error)

	InsertProduct(product Product) error
}
