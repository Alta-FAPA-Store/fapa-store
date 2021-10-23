package migration

import (
	"go-hexagonal/modules/pet"
	"go-hexagonal/modules/product"
	"go-hexagonal/modules/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.UserTable{}, &pet.PetTable{}, &product.ProductTable{}, &product.CategoryTable{}, &product.ProductCategoryTable{}, &product.ProductGalleriesTable{})
}
