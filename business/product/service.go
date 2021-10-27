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
	CategoryId  int    `validate:"required"`
}
type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) FindProductByID(id int) (*Product, error) {
	return s.repo.FindProductByID(id)
}

func (s *service) InsertProduct(insertProductSpec InsertProductSpec, createdBy string) error {
	err := validator.GetValidator().Struct(insertProductSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	product := NewProduct(
		insertProductSpec.ID,
		insertProductSpec.Name,
		insertProductSpec.Description,
		insertProductSpec.Stock,
		insertProductSpec.Price,
		insertProductSpec.CategoryId,
		createdBy,
		time.Now(),
	)

	err = s.repo.InsertProduct(product)
	if err != nil {
		return err
	}

	return nil
}
