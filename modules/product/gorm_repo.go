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
	gorm.Model
	ID                    int                     `gorm:"id"`
	ProductGalleriesTable []ProductGalleriesTable `gorm:"foreignKey:ProductTableID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProductCategory       []ProductCategoryTable  `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Name                  string                  `gorm:"name"`
	Price                 int                     `gorm:"price"`
	Description           int                     `gorm:"description"`
	Slug                  string                  `gorm:"slug"`
	Stock                 string                  `gorm:"stock"`
	ModifiedAt            time.Time               `gorm:"modified_at"`
	ModifiedBy            string                  `gorm:"modified_by"`
	Version               int                     `gorm:"version"`
}

type CategoryTable struct {
	ID         int       `gorm:"id"`
	Name       string    `gorm:"name"`
	ModifiedAt time.Time `gorm:"modified_at"`
	ModifiedBy string    `gorm:"modified_by"`
	gorm.Model
	Version int `gorm:"version"`
}

type ProductCategoryTable struct {
	ID         int           `gorm:"id"`
	CategoryID int           `gorm:"category_id"`
	ProductID  int           `gorm:"product_id"`
	Category   CategoryTable `gorm:"foreignKey:CategoryID"`
	Product    ProductTable  `gorm:"foreignKey:ProductID"`
	gorm.Model
}

type ProductGalleriesTable struct {
	ID             int       `gorm:"id"`
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

	product.ID = col.ID

	return product
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
	// err := repo.DB.Raw(`SELECT pt.id as product_id ,ct."name" as category_name,pt."name" ,pt.price, pt.slug ,pt.stock FROM product_tables pt
	// left join product_category_tables pct on pt.id = pct.product_id
	// left join category_tables ct on pct.category_id = ct.id`).Scan(&productData)

	if err != nil {
		return nil, err
	}

	product := productData.ToProduct()

	return &product, nil

}
