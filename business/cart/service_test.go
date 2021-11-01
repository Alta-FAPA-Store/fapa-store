package cart_test

import (
	"go-hexagonal/business"
	"go-hexagonal/business/cart"
	cartMock "go-hexagonal/business/cart/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id          = 1
	cart_id     = 1
	user_id     = 1
	product_id  = 1
	is_checkout = false
	quantity    = 1
)

var (
	cartService    cart.Service
	cartRepository cartMock.Repository

	cartData              cart.Cart
	insertCartData        cart.InsertCartSpec
	deleteCartDetails     cart.DeleteCartDetailsSpec
	updateCartDetailsSpec cart.UpdateCartDetailsSpec

	insertCartFailed            cart.InsertCartSpec
	deleteCartDetailsFailed     cart.DeleteCartDetailsSpec
	updateCartDetailsSpecFailed cart.UpdateCartDetailsSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindCartByUserId(t *testing.T) {
	t.Run("Expect found cart user", func(t *testing.T) {
		cartRepository.On("FindCartByUserId", mock.AnythingOfType("int")).Return(&cartData, nil).Once()

		cart, err := cartService.FindCartByUserId(user_id)

		assert.Nil(t, err)
		assert.NotNil(t, cart)

		assert.Equal(t, id, cart.Id)
		assert.Equal(t, user_id, cart.UserId)
		assert.Equal(t, is_checkout, cart.IsCheckout)
	})

	t.Run("Expect cart not found", func(t *testing.T) {
		cartRepository.On("FindCartByUserId", mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()

		user, err := cartService.FindCartByUserId(69)

		assert.NotNil(t, err)
		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestInsertCart(t *testing.T) {
	t.Run("Expect insert cart success", func(t *testing.T) {
		cartRepository.On("FindCartByUserId", mock.AnythingOfType("int")).Return(&cartData, nil).Once()
		cartRepository.On("InsertCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()

		err := cartService.InsertCart(insertCartData)

		assert.Nil(t, err)
	})
}

func TestDeleteCartDetails(t *testing.T) {
	t.Run("Expect delete cart success", func(t *testing.T) {
		cartRepository.On("DeleteCartDetails", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()

		err := cartService.DeleteCartDetails(deleteCartDetails)

		assert.Nil(t, err)
	})

	t.Run("Expect delete cart failed,cart/product id not found", func(t *testing.T) {
		cartRepository.On("DeleteCartDetails", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(business.ErrNotFound).Once()

		err := cartService.DeleteCartDetails(deleteCartDetailsFailed)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestUpdateQuantityCartDetails(t *testing.T) {

	t.Run("Expect update quantity product in cart success", func(t *testing.T) {
		cartRepository.On("UpdateQuantityCartDetails", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil).Once()

		err := cartService.UpdateQuantityCartDetails(updateCartDetailsSpec)

		assert.Nil(t, err)
	})

	t.Run("Expect update quantity product in cart failed,cart/product id not found", func(t *testing.T) {
		cartRepository.On("UpdateQuantityCartDetails", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(business.ErrNotFound).Once()

		err := cartService.UpdateQuantityCartDetails(updateCartDetailsSpecFailed)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})
}

func setup() {

	cartData = cart.NewCart(
		id,
		user_id,
		is_checkout,
		time.Now(),
	)

	insertCartData = cart.InsertCartSpec{
		UserId:    user_id,
		ProductId: product_id,
	}

	deleteCartDetails = cart.DeleteCartDetailsSpec{
		CartId:    cart_id,
		ProductId: product_id,
	}

	updateCartDetailsSpec = cart.UpdateCartDetailsSpec{
		CartId:    cart_id,
		ProductId: product_id,
		Quantity:  quantity,
	}

	insertCartFailed = cart.InsertCartSpec{
		UserId:    69,
		ProductId: 10,
	}

	deleteCartDetailsFailed = cart.DeleteCartDetailsSpec{
		CartId:    69,
		ProductId: 69,
	}

	updateCartDetailsSpecFailed = cart.UpdateCartDetailsSpec{
		CartId:    69,
		ProductId: 69,
		Quantity:  1,
	}

	cartService = cart.NewService(&cartRepository)
}
