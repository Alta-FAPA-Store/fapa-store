package product

import (
	"go-hexagonal/business"
	"go-hexagonal/util/validator"
	"time"
)

type InsertProductSpec struct {
	ID          int    `validate:"required"`
	Name        string `validate:"required"`
	Description string `validate:"required"`
	Stock       int    `validate:"required"`
	Price       int    `validate:"required"`
	Category    int
}
type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) FindProductByID(id int, userID int) (*Product, error) {
	return s.repo.FindProductByID(id, userID)
}

func (s *service) InsertProduct(insertProductSpec InsertProductSpec, createdBy string) error {
	err := validator.GetValidator().Struct(insertProductSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	product := NewProduct(
		0,
		insertProductSpec.Name,
		insertProductSpec.Description,
		insertProductSpec.Stock,
		insertProductSpec.Price,
		createdBy,
		time.Now(),
	)

	err = s.repo.InsertProduct(product)
	if err != nil {
		return err
	}

	return nil
}
