package product

type Service interface {
	//FindProductByID If data not found will return nil without error
	FindProductByID(id int, userID int) (*Product, error)
}

type Repository interface {
	//FindProductByID If data not found will return nil without error
	FindProductByID(id int, userID int) (*Product, error)
}
