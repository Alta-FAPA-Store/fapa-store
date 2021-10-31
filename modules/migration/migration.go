package migration

import (
	"go-hexagonal/modules/category"
	"go-hexagonal/modules/product"

	"go-hexagonal/modules/cart"
	"go-hexagonal/modules/transaction"

	"go-hexagonal/modules/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {

	db.AutoMigrate(&user.UserTable{}, &cart.Cart{}, &cart.CartDetails{}, &transaction.Transaction{}, &category.CategoryTable{}, &product.ProductGalleriesTable{})

}
