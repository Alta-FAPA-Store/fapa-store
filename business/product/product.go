package product

import "time"

type Product struct {
	ID, Price, Stock, Version                      int
	Name, Description, Slug, CreatedBy, ModifiedBy string
	Category                                       Category
	Photo                                          Photo
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
