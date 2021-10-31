package category_test

import (
	"go-hexagonal/business"
	"go-hexagonal/business/category"
	categoryMock "go-hexagonal/business/category/mocks"
	"time"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	name       = "elektronik"
	categoryId = 1
	creator    = "admin"
)

var (
	categoryService    category.Service
	categoryRepository categoryMock.Repository

	categoryData       category.Category
	categoryDataAll    []category.Category
	insertCategoryData category.InsertCategorySpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
func TestFindCategoryByID(t *testing.T) {
	t.Run("Expect found the category", func(t *testing.T) {
		categoryRepository.On("FindCategoryByID", mock.AnythingOfType("int")).Return(&categoryData, nil).Once()

		category, err := categoryService.FindCategoryByID(categoryId)

		assert.Nil(t, err)
		assert.NotNil(t, category)
		assert.Equal(t, name, category.CategoryName)

	})

	t.Run("Expect category not found", func(t *testing.T) {
		categoryRepository.On("FindCategoryByID", mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()

		category, err := categoryService.FindCategoryByID(categoryId)

		assert.NotNil(t, err)

		assert.Nil(t, category)

		assert.Equal(t, err, business.ErrNotFound)
	})

}

func TestFindAllCategory(t *testing.T) {
	t.Run("Expect found the category", func(t *testing.T) {
		categoryRepository.On("FindAllCategory", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(&categoryDataAll, nil).Once()

		category, err := categoryService.FindAllCategory(1, 10)

		assert.Nil(t, err)

		assert.NotNil(t, category)

		assert.Equal(t, categoryId, category[0].CategoryID)
		assert.Equal(t, name, category[0].CategoryName)

	})

	// t.Run("Expect category not found", func(t *testing.T) {
	// 	// categoryRepository.On("FindAllCategory", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
	// 	categoryRepository.On("FindAllCategory", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

	// 	category, err := categoryService.FindAllCategory(1, 10, "elektronik", name)

	// 	assert.NotNil(t, err)

	// 	assert.Nil(t, category)

	// 	assert.Equal(t, err, business.ErrNotFound)
	// })
	// t.Run("Expect category not found", func(t *testing.T) {
	// 	// categoryRepository.On("FindAllCategory", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
	// 	categoryRepository.On("FindAllCategory", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

	// 	category, err := categoryService.FindAllCategory(1, 10, "elektronik", name)

	// 	assert.NotNil(t, err)

	// 	assert.Nil(t, category)

	// 	assert.Equal(t, err, business.ErrNotFound)
	// })
}

func TestInsertCategory(t *testing.T) {
	t.Run("Expect insert category success", func(t *testing.T) {
		categoryRepository.On("InsertCategory", mock.AnythingOfType("category.Category")).Return(nil).Once()

		err := categoryService.InsertCategory(insertCategoryData, creator)

		assert.Nil(t, err)

	})

	t.Run("Expect insert category not found", func(t *testing.T) {
		categoryRepository.On("InsertCategory", mock.AnythingOfType("category.Category")).Return(business.ErrInternalServerError).Once()

		err := categoryService.InsertCategory(insertCategoryData, creator)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expect insert category invalid spec", func(t *testing.T) {
		categoryRepository.On("InsertCategory", mock.AnythingOfType("category.Category")).Return(business.ErrInvalidSpec).Once()

		insertCategoryData.Name = ""
		err := categoryService.InsertCategory(insertCategoryData, creator)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInvalidSpec)
	})
}

func TestUpdateCategoryById(t *testing.T) {
	t.Run("Expect update category success", func(t *testing.T) {
		categoryRepository.On("FindCategoryByID", mock.AnythingOfType("int")).Return(&categoryData, nil).Once()
		categoryRepository.On("UpdateCategory", mock.AnythingOfType("category.Category")).Return(nil).Once()

		err := categoryService.UpdateCategory(categoryId, category.UpdateCategorySpec{})

		assert.Nil(t, err)

	})

	t.Run("Expect update category failed", func(t *testing.T) {
		categoryRepository.On("FindCategoryByID", mock.AnythingOfType("int")).Return(&categoryData, nil).Once()
		categoryRepository.On("UpdateCategory", mock.AnythingOfType("category.Category")).Return(business.ErrInternalServerError).Once()

		err := categoryService.UpdateCategory(categoryId, category.UpdateCategorySpec{})

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expect category not found", func(t *testing.T) {
		categoryRepository.On("FindCategoryByID", mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()
		categoryRepository.On("UpdateCategory", mock.AnythingOfType("category.Category")).Return(business.ErrInternalServerError).Once()

		err := categoryService.UpdateCategory(categoryId, category.UpdateCategorySpec{})

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrNotFound)
	})

	t.Run("Expect category failed", func(t *testing.T) {
		categoryRepository.On("FindCategoryByID", mock.AnythingOfType("int")).Return(nil, business.ErrInternalServerError).Once()
		categoryRepository.On("UpdateCategory", mock.AnythingOfType("category.Category")).Return(business.ErrInternalServerError).Once()

		err := categoryService.UpdateCategory(categoryId, category.UpdateCategorySpec{})

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})

}

func TestDeleteCategoryByID(t *testing.T) {
	t.Run("Expect delete the category", func(t *testing.T) {
		categoryRepository.On("DeleteCategory", mock.AnythingOfType("int")).Return(nil).Once()

		err := categoryService.DeleteCategory(categoryId)

		assert.Nil(t, err)

	})

	t.Run("Expect category not found", func(t *testing.T) {
		categoryRepository.On("DeleteCategory", mock.AnythingOfType("int")).Return(business.ErrInternalServerError).Once()

		err := categoryService.DeleteCategory(categoryId)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})

}

func setup() {

	categoryData = category.NewCategory(
		name,
		time.Now(),
	)
	categoryData.CategoryID = categoryId

	categoryDataAll = append(categoryDataAll, categoryData)

	insertCategoryData = category.InsertCategorySpec{
		Name: name,
	}

	categoryService = category.NewService(&categoryRepository)
}
