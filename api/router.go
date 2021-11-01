package api

import (
	"go-hexagonal/api/middleware"
	"go-hexagonal/api/v1/auth"

	"go-hexagonal/api/v1/category"
	"go-hexagonal/api/v1/product"

	"go-hexagonal/api/v1/cart"
	"go-hexagonal/api/v1/transaction"

	"go-hexagonal/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

//RegisterPath Register all API with routing path

func RegisterPath(e *echo.Echo, authController *auth.Controller, userController *user.Controller, cartController *cart.Controller, transactionController *transaction.Controller, productController *product.Controller, categoryController *category.Controller) {
	if authController == nil || userController == nil || cartController == nil || transactionController == nil || productController == nil || categoryController == nil {

		panic("Controller parameter cannot be nil")
	}
	//authentication with Versioning endpoint
	authV1 := e.Group("v1/auth")
	authV1.POST("/login", authController.Login)
	authV1.POST("/register", authController.Register)

	//user with Versioning endpoint
	userV1 := e.Group("v1/users")
	// userV1.GET("/:id", userController.FindUserByID)
	userV1.POST("", userController.InsertUser)
	userV1.GET("/:username", userController.FindUserByUsername)
	userV1.GET("", userController.FindAllUser)
	userV1.PUT("/:username", userController.UpdateUser)

	//product with versioning endpoint
	productV1 := e.Group("v1/products")
	productV1.Use(middleware.JWTMiddleware())
	productV1.GET("/:id", productController.FindProductByID)
	productV1.GET("", productController.FindAllProduct)
	productV1.POST("", productController.InsertProduct)
	productV1.PUT("/:id", productController.UpdateProduct)
	productV1.DELETE("/:id", productController.DeleteProduct)

	categoryV1 := e.Group("v1/category")
	categoryV1.Use(middleware.JWTMiddleware())
	categoryV1.GET("/:id", categoryController.FindCategoryByID)
	categoryV1.GET("", categoryController.FindAllCategory)
	categoryV1.POST("", categoryController.InsertCategory)
	categoryV1.PUT("/:id", categoryController.UpdateCategory)
	categoryV1.DELETE("/:id", categoryController.DeleteCategory)

	// Cart with versioning endpoint
	cartV1 := e.Group("v1/cart")
	cartV1.Use(middleware.JWTMiddleware())
	cartV1.GET("/:user_id", cartController.FindCartByUserId)
	cartV1.POST("", cartController.InsertCart)
	cartV1.PUT("", cartController.UpdateQuantityCartDetails)
	cartV1.DELETE("", cartController.DeleteCartDetails)

	// Transaction with versioning endpoint
	transactionV1 := e.Group("v1/transaction")
	transactionV1.Use(middleware.JWTMiddleware())
	transactionV1.GET("", transactionController.GetAllTransaction)
	transactionV1.GET("/:transaction_id", transactionController.GetTransactionDetails)
	transactionV1.POST("", transactionController.CreateTransaction)
	transactionV1.PUT("/:transaction_id", transactionController.UpdateTransaction)
	transactionV1.DELETE("/:transaction_id", transactionController.DeleteTransaction)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
