package api

import (
	"go-hexagonal/api/middleware"
	"go-hexagonal/api/v1/auth"
	"go-hexagonal/api/v1/pet"
	"go-hexagonal/api/v1/product"
	"go-hexagonal/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

//RegisterPath Register all API with routing path
func RegisterPath(e *echo.Echo, authController *auth.Controller, userController *user.Controller, petController *pet.Controller, productController *product.Controller) {
	if authController == nil || userController == nil || petController == nil {
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

	//product with versioning endpoint
	productV1 := e.Group("v1/products")
	// productV1.GET("/:id", petController.FindPetByID)
	// productV1.GET("", petController.FindAllPet)
	productV1.POST("", productController.InsertProduct)
	// productV1.PUT("/:id", petController.UpdatePet)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
