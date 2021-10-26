package product

import "time"

type Product struct {
	ID, Price, Stock, Version                      int
	Name, Description, Slug, CreatedBy, ModifiedBy string
	Category                                       []Category
	Photo                                          []Photo
	CreatedAt                                      time.Time
	ModifiedAt                                     time.Time
}

type Category struct {
	CategoryID   int
	CategoryName string
}

type Photo struct {
	URL interface{}
}

func NewProduct(
	id int,
	name string,
	description string,
	stock int,
	price int,
	creator string,
	createdAt time.Time) Product {

	return Product{
		ID:          id,
		Price:       price,
		Stock:       stock,
		Version:     1,
		Name:        name,
		Description: description,
		Slug:        "",
		CreatedBy:   creator,
		ModifiedBy:  creator,
		Category:    []Category{},
		Photo:       []Photo{},
		CreatedAt:   createdAt,
		ModifiedAt:  createdAt,
	}
}
