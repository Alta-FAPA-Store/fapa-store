package product

import (
	"go-hexagonal/business/product"
	"time"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}
type ProductTable struct {
	ID          int       `gorm:"id"`
	Name        string    `gorm:"name"`
	Price       int       `gorm:"price"`
	Description string    `gorm:"description"`
	Slug        string    `gorm:"slug"`
	Stock       int       `gorm:"stock"`
	CategoryID  int       `gorm:"category_id"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
	DeletedAt   gorm.DeletedAt
}

type ProductGalleriesTable struct {
	Id         int       `gorm:"id"`
	ProductID  int       `gorm:"product_id"`
	URL        string    `gorm:"url"`
	IsFeatured bool      `gorm:"is_featured"`
	CreatedAt  time.Time `gorm:"created_at"`
	UpdatedAt  time.Time `gorm:"updated_at"`
	DeletedAt  gorm.DeletedAt
}

func (col *ProductTable) ToProduct() product.Product {
	var product product.Product

	product.ID = col.ID
	product.Name = col.Name
	product.Description = col.Description
	product.Price = col.Price
	product.Slug = col.Slug
	product.Stock = col.Stock
	product.CategoryID = col.CategoryID

	return product
}

func newProductTable(product product.Product) *ProductTable {

	return &ProductTable{
		product.ID,
		product.Name,
		product.Price,
		product.Description,
		product.Slug,
		product.Stock,
		product.CategoryID,
		product.CreatedAt,
		product.UpdatedAt,
		product.DeletedAt,
	}

}

//NewGormDBRepository Generate Gorm DB pet repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (repo *GormRepository) FindProductByID(id int) (*product.Product, error) {
	var productData ProductTable

	err := repo.DB.Where("id = ?", id).First(&productData).Error

	if err != nil {
		return nil, err
	}

	product := productData.ToProduct()

	return &product, nil

}

func (repo *GormRepository) FindAllProduct(skip int, rowPerPage int, categoryParam, nameParam string) ([]product.Product, error) {

	var products []ProductTable

	var err error

	err = repo.DB.Offset(skip).Limit(rowPerPage).Find(&products).Error

	if categoryParam != "" {
		err = repo.DB.Where("category_id = ?", categoryParam).Offset(skip).Limit(rowPerPage).Find(&products).Error
	}
	if nameParam != "" {
		err = repo.DB.Where("name = ?", nameParam).Offset(skip).Limit(rowPerPage).Find(&products).Error
	}

	if nameParam != "" && categoryParam != "" {
		err = repo.DB.Where("category_id = ?", categoryParam).Where("name = ?", nameParam).Offset(skip).Limit(rowPerPage).Find(&products).Error
	}

	if err != nil {
		return nil, err
	}

	var result []product.Product
	for _, value := range products {
		result = append(result, value.ToProduct())
	}

	return result, nil
}

func (repo *GormRepository) InsertProduct(product product.Product) error {

	productData := newProductTable(product)

	err := repo.DB.Create(productData).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *GormRepository) UpdateProduct(product product.Product) error {

	productData := newProductTable(product)

	err := repo.DB.Model(&productData).Where("id = ?", product.ID).Updates(
		ProductTable{
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Stock:       product.Stock,
			CategoryID:  product.CategoryID,
			UpdatedAt:   product.UpdatedAt,
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) DeleteProduct(id int) error {
	var product ProductTable
	err := repo.DB.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		return err
	}

	return nil
}
