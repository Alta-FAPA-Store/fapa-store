package category

import (
	"go-hexagonal/business"
	"go-hexagonal/util/validator"
	"time"
)

type InsertCategorySpec struct {
	Name string `validate:"required"`
}

type UpdateCategorySpec struct {
	Name string `validate:"required"`
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) FindCategoryByID(id int) (*Category, error) {
	category, err := s.repo.FindCategoryByID(id)

	if err != nil {
		return nil, business.ErrNotFound
	}

	return category, nil
}

func (s *service) FindAllCategory(skip int, rowPerPage int) ([]Category, error) {

	category, err := s.repo.FindAllCategory(skip, rowPerPage)
	if err != nil {
		return []Category{}, err
	}

	return category, err
}

func (s *service) InsertCategory(insertCategorySpec InsertCategorySpec, createdBy string) error {
	err := validator.GetValidator().Struct(insertCategorySpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	category := NewCategory(
		insertCategorySpec.Name,
		time.Now(),
	)

	err = s.repo.InsertCategory(category)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateCategory(id int, updateCategorySpec UpdateCategorySpec) error {

	category, err := s.repo.FindCategoryByID(id)
	if err != nil {
		return err
	} else if category == nil {
		return business.ErrNotFound
	}

	category.UpdatedAt = time.Now()
	category.CategoryName = updateCategorySpec.Name

	return s.repo.UpdateCategory(*category)
}

func (s *service) DeleteCategory(id int) error {
	err := s.repo.DeleteCategory(id)

	if err != nil {
		return err
	}

	return nil
}
