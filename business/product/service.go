package product

import (
	"go-hexagonal/business"
	"go-hexagonal/util/validator"
	"time"
)

type InsertProductSpec struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
	Stock       int    `validate:"required"`
	Price       int    `validate:"required"`
	CategoryId  int    `validate:"required"`
}

type UpdateProductSpec struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
	Stock       int    `validate:"number"`
	Price       int    `validate:"number"`
	CategoryId  int    `validate:"number"`
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

	product, err := s.repo.FindProductByID(id)

	if err != nil {
		return nil, business.ErrNotFound
	}

	return product, nil
}

func (s *service) FindAllProduct(skip int, rowPerPage int, categoryParam, nameParam string) ([]Product, error) {

	product, err := s.repo.FindAllProduct(skip, rowPerPage, categoryParam, nameParam)
	if err != nil {
		return nil, err
	}

	return product, err
}

func (s *service) InsertProduct(insertProductSpec InsertProductSpec, createdBy string) error {
	err := validator.GetValidator().Struct(insertProductSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	product := NewProduct(
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

func (s *service) UpdateProduct(id int, updateProductSpec UpdateProductSpec) error {

	product, err := s.repo.FindProductByID(id)
	if err != nil {
		return err
	} else if product == nil {
		return business.ErrNotFound
	}

	product.UpdatedAt = time.Now()
	product.Name = updateProductSpec.Name
	product.Description = updateProductSpec.Description
	product.Stock = updateProductSpec.Stock
	product.CategoryID = updateProductSpec.CategoryId
	product.Price = updateProductSpec.Price

	// modifiedPet := product.ModifyPet(name, time.Now(), modifiedBy)

	return s.repo.UpdateProduct(*product)
}

func (s *service) DeleteProduct(id int) error {
	err := s.repo.DeleteProduct(id)

	if err != nil {
		return err
	}

	return nil
}
