package migration

import (
	"go-hexagonal/modules/category"
	"go-hexagonal/modules/pet"
	"go-hexagonal/modules/product"
	"go-hexagonal/modules/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.UserTable{}, &pet.PetTable{}, &product.ProductTable{}, &category.CategoryTable{}, &product.ProductGalleriesTable{})
}
