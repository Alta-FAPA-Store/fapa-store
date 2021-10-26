package api

import (
	"go-hexagonal/api/middleware"
	"go-hexagonal/api/v1/auth"
	"go-hexagonal/api/v1/cart"
	"go-hexagonal/api/v1/pet"
	"go-hexagonal/api/v1/transaction"
	"go-hexagonal/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

//RegisterPath Register all API with routing path
func RegisterPath(e *echo.Echo, authController *auth.Controller, userController *user.Controller, petController *pet.Controller, cartController *cart.Controller, transactionController *transaction.Controller) {
	if authController == nil || userController == nil || petController == nil || cartController == nil || transactionController == nil {
		panic("Controller parameter cannot be nil")
	}

	//authentication with Versioning endpoint
	authV1 := e.Group("v1/auth")
	authV1.POST("/login", authController.Login)

	//user with Versioning endpoint
	userV1 := e.Group("v1/users")
	// userV1.GET("/:id", userController.FindUserByID)
	userV1.Use(middleware.JWTMiddleware())
	userV1.GET("/:username", userController.FindUserByUsername)
	userV1.GET("", userController.FindAllUser)
	userV1.POST("", userController.InsertUser)
	userV1.PUT("/:username", userController.UpdateUser)

	//pet with Versioning endpoint
	petV1 := e.Group("v1/pets")
	petV1.GET("/:id", petController.FindPetByID)
	petV1.GET("", petController.FindAllPet)
	petV1.POST("", petController.InsertPet)
	petV1.PUT("/:id", petController.UpdatePet)

	// Cart with versioning endpoint
	cartV1 := e.Group("v1/cart")
	cartV1.GET("/:user_id", cartController.FindCartByUserId)
	cartV1.POST("", cartController.InsertCart)
	cartV1.PUT("", cartController.UpdateQuantityCartDetails)
	cartV1.DELETE("", cartController.DeleteCartDetails)

	// Transaction with versioning endpoint
	transactionV1 := e.Group("v1/transaction")
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
