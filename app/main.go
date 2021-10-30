package main

import (
	"context"
	"fmt"
	"go-hexagonal/config"

	"go-hexagonal/api"
	userController "go-hexagonal/api/v1/user"
	userService "go-hexagonal/business/user"
	migration "go-hexagonal/modules/migration"
	userRepository "go-hexagonal/modules/user"

	productController "go-hexagonal/api/v1/product"
	productService "go-hexagonal/business/product"
	productRepository "go-hexagonal/modules/product"

	categoryController "go-hexagonal/api/v1/category"
	categoryService "go-hexagonal/business/category"
	categoryRepository "go-hexagonal/modules/category"

	authController "go-hexagonal/api/v1/auth"
	authService "go-hexagonal/business/auth"

	cartController "go-hexagonal/api/v1/cart"
	cartService "go-hexagonal/business/cart"
	cartRepository "go-hexagonal/modules/cart"

	transactionController "go-hexagonal/api/v1/transaction"
	transactionService "go-hexagonal/business/transaction"
	transactionRepository "go-hexagonal/modules/transaction"

	"os"
	"os/signal"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	// "gorm.io/driver/mysql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newDatabaseConnection(config *config.AppConfig) *gorm.DB {

	// configDB := map[string]string{
	// 	"DB_Username": os.Getenv("GOHEXAGONAL_DB_USERNAME"),
	// 	"DB_Password": os.Getenv("GOHEXAGONAL_DB_PASSWORD"),
	// 	"DB_Port":     os.Getenv("GOHEXAGONAL_DB_PORT"),
	// 	"DB_Host":     os.Getenv("GOHEXAGONAL_DB_ADDRESS"),
	// 	"DB_Name":     os.Getenv("GOHEXAGONAL_DB_NAME"),
	// }

	connectionString := "host=localhost user=postgres password=root dbname=fapa-store port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
	// 	configDB["DB_Host"],
	// 	configDB["DB_Username"],
	// 	configDB["DB_Password"],
	// 	configDB["DB_Name"],
	// 	configDB["DB_Port"])

	// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	configDB["DB_Username"],
	// 	configDB["DB_Password"],
	// 	configDB["DB_Host"],
	// 	configDB["DB_Port"],
	// 	configDB["DB_Name"])

	db, e := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	migration.InitMigrate(db)

	return db
}

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	dbConnection := newDatabaseConnection(config)
	//initiate user repository
	userRepo := userRepository.NewGormDBRepository(dbConnection)

	//initiate user service
	userService := userService.NewService(userRepo)

	//initiate user controller
	userController := userController.NewController(userService)

	//initiate pet repository
	// petRepo := petRepository.NewGormDBRepository(dbConnection)

	// //initiate pet service
	// petService := petService.NewService(petRepo)

	// //initiate pet controller
	// petController := petController.NewController(petService)

	//initiate auth service
	authService := authService.NewService(userService)

	//initiate auth controller
	authController := authController.NewController(authService)

	productRepo := productRepository.NewGormDBRepository(dbConnection)
	productService := productService.NewService(productRepo)
	productController := productController.NewController(productService)

	categoryRepo := categoryRepository.NewGormDBRepository(dbConnection)
	categoryService := categoryService.NewService(categoryRepo)
	categoryController := categoryController.NewController(categoryService)

	//initiate cart repository
	cartRepo := cartRepository.NewGormDBRepository(dbConnection)

	//initiate cart service
	cartService := cartService.NewService(cartRepo)

	//initiate cart controller
	cartController := cartController.NewController(cartService)

	//initiate transaction repository
	transactionRepo := transactionRepository.NewGormDBRepository(dbConnection)

	//initiate transaction service
	transactionService := transactionService.NewService(transactionRepo)

	//initiate transaction controller
	transactionController := transactionController.NewController(transactionService)

	//create echo http
	e := echo.New()

	//register API path and handler

	api.RegisterPath(e, authController, userController, cartController, transactionController, productController, categoryController)

	// run server
	go func() {
		address := fmt.Sprintf("localhost:%d", config.AppPort)
		if err := e.Start(address); err != nil {
			fmt.Println(err)
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
