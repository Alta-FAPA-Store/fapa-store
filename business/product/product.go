package product

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID, Price, Stock, Version          int
	Name, Description, Slug, CreatedBy string
	CategoryID                         int
	CreatedAt                          time.Time
	UpdatedAt                          time.Time
	DeletedAt                          gorm.DeletedAt
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
	categoryId int,
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
		CategoryID:  categoryId,
		CreatedBy:   creator,
		CreatedAt:   createdAt,
		UpdatedAt:   createdAt,
	}
}
