package cart

type Service interface {
	FindCartByUserId(userId int) (*Cart, error)
	InsertCart(insertCartSpec InsertCartSpec) error
	DeleteCartDetails(deleteCartDetails DeleteCartDetailsSpec) error
}

type Repository interface {
	InsertCart(cart Cart) (int, error)
	InsertCartDetails(cartDetails CartDetails) error
	FindCartByUserId(userId int) (*Cart, error)
	CheckCartProduct(cartId int, productId int) (bool, error)
	UpdateCartDetailsProduct(cartId int, productId int) error
	DeleteCartDetails(cartId int, productId int) error
}
