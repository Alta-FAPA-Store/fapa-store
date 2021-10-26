package product

import (
	"go-hexagonal/business/product"
	"time"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

// type ProductTable struct {
// 	gorm.Model
// 	Name                  string                  `gorm:"name"`
// 	Price                 int                     `gorm:"price"`
// 	Description           int                     `gorm:"description"`
// 	Slug                  string                  `gorm:"slug"`
// 	Stock                 string                  `gorm:"stock"`
// 	ProductGalleriesTable []ProductGalleriesTable `gorm:"foreignKey:ProductTableID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
// 	ProductCategory       ProductCategoryTable    `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
// 	ModifiedAt            time.Time               `gorm:"modified_at"`
// 	ModifiedBy            string                  `gorm:"modified_by"`
// 	Version               int                     `gorm:"version"`
// }
type ProductTable struct {
	gorm.Model
	Name        string `gorm:"name"`
	Price       int    `gorm:"price"`
	Description string `gorm:"description"`
	Slug        string `gorm:"slug"`
	Stock       int    `gorm:"stock"`
	Version     int    `gorm:"version"`
}

type CategoryTable struct {
	Name            string                 `gorm:"name"`
	ModifiedAt      time.Time              `gorm:"modified_at"`
	ModifiedBy      string                 `gorm:"modified_by"`
	ProductCategory []ProductCategoryTable `gorm:"foreignKey:ID"`
	gorm.Model
	Version int `gorm:"version"`
}

type ProductCategoryTable struct {
	CategoryID int `gorm:"category_id"`
	ProductID  int `gorm:"product_id"`
	gorm.Model
}

type ProductGalleriesTable struct {
	ProductTableID int       `gorm:"product_id"`
	URL            string    `gorm:"url"`
	IsFeatured     string    `gorm:"is_featured"`
	ModifiedAt     time.Time `gorm:"modified_at"`
	ModifiedBy     string    `gorm:"modified_by"`
	gorm.Model
	Version int `gorm:"version"`
}

func (col *ProductTable) ToProduct() product.Product {
	var product product.Product

	product.Name = col.Name

	return product
}

func newProductTable(product product.Product) *ProductTable {

	return &ProductTable{
		gorm.Model{
			ID:        uint(product.ID),
			CreatedAt: product.CreatedAt,
			UpdatedAt: time.Now(),
		},
		product.Name,
		product.Price,
		product.Description,
		product.Slug,
		product.Stock,
		product.Version,
	}

}

//NewGormDBRepository Generate Gorm DB pet repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (repo *GormRepository) FindProductByID(id, userID int) (*product.Product, error) {
	var productData ProductTable

	err := repo.DB.Where("id = ?", id).Where("user_id = ?", userID).First(&productData).Error

	if err != nil {
		return nil, err
	}

	product := productData.ToProduct()

	return &product, nil

}

func (repo *GormRepository) InsertProduct(product product.Product) error {

	productData := newProductTable(product)
	productData.ID = 0

	err := repo.DB.Create(productData).Error
	if err != nil {
		return err
	}
	return nil
}
