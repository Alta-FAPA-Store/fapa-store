package product

import (
	"fmt"
	"go-hexagonal/api/common"
	"go-hexagonal/api/v1/product/request"
	"go-hexagonal/api/v1/product/response"
	"go-hexagonal/business/product"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service product.Service
}

//NewController Construct item API controller
func NewController(service product.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) FindProductByID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	id, _ := strconv.Atoi(c.Param("id"))

	product, err := controller.service.FindProductByID(id)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetProductResponse(*product)

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) InsertProduct(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	//use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
	//MUST CONVERT TO FLOAT64, OTHERWISE ERROR (not _ok_)!
	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	fmt.Println(userID)

	insertProductRequest := new(request.InsertProductRequest)
	if err := c.Bind(insertProductRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.InsertProduct(*insertProductRequest.ToUpsertProductSpec(), "creator")
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
