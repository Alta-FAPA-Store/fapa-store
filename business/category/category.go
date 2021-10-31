package category

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	CategoryID   int
	CategoryName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

func NewCategory(
	name string,
	createdAt time.Time) Category {

	return Category{
		CategoryName: name,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
	}
}
