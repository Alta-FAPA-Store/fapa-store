package product_test

import (
	"go-hexagonal/business"
	"go-hexagonal/business/product"
	productMock "go-hexagonal/business/product/mocks"
	"time"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id          = 1
	price       = 20000
	stock       = 3
	name        = "Lenove"
	description = "RAM 8 GB"
	slug        = ""
	categoryId  = 1
	creator     = "admin"
)

var (
	productService    product.Service
	productRepository productMock.Repository

	productData       product.Product
	productDataAll    []product.Product
	insertProductData product.InsertProductSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
func TestFindProductByID(t *testing.T) {
	t.Run("Expect found the product", func(t *testing.T) {
		productRepository.On("FindProductByID", mock.AnythingOfType("int")).Return(&productData, nil).Once()

		product, err := productService.FindProductByID(id)

		assert.Nil(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, id, product.ID)
		assert.Equal(t, name, product.Name)
		assert.Equal(t, description, product.Description)
		assert.Equal(t, stock, product.Stock)

	})

	t.Run("Expect product not found", func(t *testing.T) {
		productRepository.On("FindProductByID", mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()

		product, err := productService.FindProductByID(id)

		assert.NotNil(t, err)

		assert.Nil(t, product)

		assert.Equal(t, err, business.ErrNotFound)
	})

}

func TestFindAllProduct(t *testing.T) {
	t.Run("Expect found the product", func(t *testing.T) {
		productRepository.On("FindAllProduct", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&productDataAll, nil).Once()

		product, err := productService.FindAllProduct(1, 10, "elektronik", name)

		assert.Nil(t, err)

		assert.NotNil(t, product)

		assert.Equal(t, id, product[0].ID)
		assert.Equal(t, name, product[0].Name)
		assert.Equal(t, description, product[0].Description)
		assert.Equal(t, stock, product[0].Stock)

	})

	// t.Run("Expect product not found", func(t *testing.T) {
	// 	// productRepository.On("FindAllProduct", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
	// 	productRepository.On("FindAllProduct", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

	// 	product, err := productService.FindAllProduct(1, 10, "elektronik", name)

	// 	assert.NotNil(t, err)

	// 	assert.Nil(t, product)

	// 	assert.Equal(t, err, business.ErrNotFound)
	// })
	// t.Run("Expect product not found", func(t *testing.T) {
	// 	// productRepository.On("FindAllProduct", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
	// 	productRepository.On("FindAllProduct", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

	// 	product, err := productService.FindAllProduct(1, 10, "elektronik", name)

	// 	assert.NotNil(t, err)

	// 	assert.Nil(t, product)

	// 	assert.Equal(t, err, business.ErrNotFound)
	// })
}

func TestInsertProduct(t *testing.T) {
	t.Run("Expect insert product success", func(t *testing.T) {
		productRepository.On("InsertProduct", mock.AnythingOfType("product.Product")).Return(nil).Once()

		err := productService.InsertProduct(insertProductData, creator)

		assert.Nil(t, err)

	})

	t.Run("Expect insert product not found", func(t *testing.T) {
		productRepository.On("InsertProduct", mock.AnythingOfType("product.Product")).Return(business.ErrInternalServerError).Once()

		err := productService.InsertProduct(insertProductData, creator)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expect insert product invalid spec", func(t *testing.T) {
		productRepository.On("InsertProduct", mock.AnythingOfType("product.Product")).Return(business.ErrInvalidSpec).Once()

		insertProductData.Name = ""
		err := productService.InsertProduct(insertProductData, creator)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInvalidSpec)
	})
}

func TestUpdateProductById(t *testing.T) {
	t.Run("Expect update product success", func(t *testing.T) {
		productRepository.On("FindProductByID", mock.AnythingOfType("int")).Return(&productData, nil).Once()
		productRepository.On("UpdateProduct", mock.AnythingOfType("product.Product")).Return(nil).Once()

		err := productService.UpdateProduct(id, product.UpdateProductSpec{})

		assert.Nil(t, err)

	})

	t.Run("Expect update product failed", func(t *testing.T) {
		productRepository.On("FindProductByID", mock.AnythingOfType("int")).Return(&productData, nil).Once()
		productRepository.On("UpdateProduct", mock.AnythingOfType("product.Product")).Return(business.ErrInternalServerError).Once()

		err := productService.UpdateProduct(id, product.UpdateProductSpec{})

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expect product not found", func(t *testing.T) {
		productRepository.On("FindProductByID", mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()
		productRepository.On("UpdateProduct", mock.AnythingOfType("product.Product")).Return(business.ErrInternalServerError).Once()

		err := productService.UpdateProduct(id, product.UpdateProductSpec{})

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrNotFound)
	})

	t.Run("Expect product failed", func(t *testing.T) {
		productRepository.On("FindProductByID", mock.AnythingOfType("int")).Return(nil, business.ErrInternalServerError).Once()
		productRepository.On("UpdateProduct", mock.AnythingOfType("product.Product")).Return(business.ErrInternalServerError).Once()

		err := productService.UpdateProduct(id, product.UpdateProductSpec{})

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})

}

func TestDeleteProductByID(t *testing.T) {
	t.Run("Expect delete the product", func(t *testing.T) {
		productRepository.On("DeleteProduct", mock.AnythingOfType("int")).Return(nil).Once()

		err := productService.DeleteProduct(id)

		assert.Nil(t, err)

	})

	t.Run("Expect product not found", func(t *testing.T) {
		productRepository.On("DeleteProduct", mock.AnythingOfType("int")).Return(business.ErrInternalServerError).Once()

		err := productService.DeleteProduct(id)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})

}

func setup() {

	productData = product.NewProduct(
		name,
		description,
		stock,
		price,
		categoryId,
		creator,
		time.Now(),
	)
	productData.ID = id

	productDataAll = append(productDataAll, productData)

	insertProductData = product.InsertProductSpec{
		Name:        name,
		Description: description,
		Stock:       stock,
		Price:       price,
		CategoryId:  categoryId,
	}

	productService = product.NewService(&productRepository)
}
