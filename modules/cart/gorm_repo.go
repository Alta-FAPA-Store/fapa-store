package cart

import (
	"go-hexagonal/business/cart"
	"time"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

type Cart struct {
	Id         int       `gorm:"id"`
	UserId     int       `gorm:"user_id"`
	IsCheckout bool      `gorm:"is_checkout"`
	CreatedAt  time.Time `gorm:"created_at"`
	UpdatedAt  time.Time `gorm:"updated_at"`
	DeletedAt  gorm.DeletedAt
}

type CartDetails struct {
	Id        int       `gorm:"id"`
	CartId    int       `gorm:"cart_id"`
	ProductId int       `gorm:"product_id"`
	Quantity  int       `gorm:"quantity"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt
}

type User struct {
	Id int
}

func newCart(cart cart.Cart) *Cart {
	return &Cart{
		cart.Id,
		cart.UserId,
		cart.IsCheckout,
		cart.CreatedAt,
		cart.UpdatedAt,
		cart.DeletedAt,
	}
}

func newCartDetails(cartDetails cart.CartDetails) *CartDetails {
	return &CartDetails{
		cartDetails.Id,
		cartDetails.CartId,
		cartDetails.ProductId,
		cartDetails.Quantity,
		cartDetails.CreatedAt,
		cartDetails.UpdatedAt,
		cartDetails.DeletedAt,
	}
}

func (col *Cart) ToCart(cartDetails []cart.CartDetailsWithProduct) cart.Cart {
	var cart cart.Cart

	cart.Id = col.Id
	cart.UserId = col.UserId
	cart.IsCheckout = col.IsCheckout
	cart.CreatedAt = col.CreatedAt
	cart.UpdatedAt = col.UpdatedAt
	cart.Details = cartDetails

	return cart
}

//NewGormDBRepository Generate Gorm DB cart repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (repo *GormRepository) InsertCart(cart cart.Cart) (int, error) {
	cartData := newCart(cart)
	cartData.Id = 0

	result := repo.DB.Create(cartData)
	err := result.Error

	if err != nil {
		return 0, err
	}

	return cartData.Id, nil
}

func (repo *GormRepository) InsertCartDetails(cartDetail cart.CartDetails) error {
	cartDetailsData := newCartDetails(cartDetail)

	err := repo.DB.Create(cartDetailsData).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) FindCartByUserId(userId int) (*cart.Cart, error) {
	var cartData Cart
	var userData User

	findUser := repo.DB.Table("user_tables").Select("id").Where("id = ?", userId).First(&userData)

	if findUser.RowsAffected == 0 {
		return nil, findUser.Error
	}

	findCart := repo.DB.Where("user_id = ?", userId).Where("is_checkout = ?", false).First(&cartData)

	if findCart.RowsAffected == 0 {
		return nil, nil
	}

	var results []cart.CartDetailsWithProduct
	repo.DB.Table("cart_details").Select("cart_details.id, product_tables.name, product_tables.price, cart_details.quantity").Joins("JOIN product_tables ON cart_details.product_id = product_tables.id").Where("cart_details.cart_id = ?", cartData.Id).Where("cart_details.deleted_at IS NULL").Scan(&results)

	cart := cartData.ToCart(results)

	return &cart, nil
}

func (repo *GormRepository) CheckCartProduct(cartId int, productId int) (bool, error) {
	var cartDetailsData CartDetails

	findCartProduct := repo.DB.Where("cart_id = ?", cartId).Where("product_id = ?", productId).First(&cartDetailsData)

	if findCartProduct.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (repo *GormRepository) UpdateCartDetailsProduct(cartId int, productId int) error {
	var cartDetailsData CartDetails

	err := repo.DB.Model(&cartDetailsData).Where("cart_id = ?", cartId).Where("product_id = ?", productId).Updates(map[string]interface{}{"quantity": gorm.Expr("quantity + ?", 1), "updated_at": time.Now()}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) DeleteCartDetails(cartId int, productId int) error {
	var cartDetail CartDetails
	err := repo.DB.Where("cart_id = ?", cartId).Where("product_id = ?", productId).Delete(&cartDetail).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) UpdateQuantityCartDetails(cartId int, productId int, quantity int) error {
	err := repo.DB.Model(&CartDetails{}).Where("cart_id = ?", cartId).Where("product_id = ?", productId).Update("quantity", quantity).Error

	if err != nil {
		return err
	}

	return nil
}
