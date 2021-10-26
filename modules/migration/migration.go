package migration

import (
	"go-hexagonal/modules/cart"
	"go-hexagonal/modules/pet"
	"go-hexagonal/modules/transaction"
	"go-hexagonal/modules/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.UserTable{}, &pet.PetTable{}, &cart.Cart{}, &cart.CartDetails{}, &transaction.Transaction{})
}
