package product

import (
	"go-hexagonal/api/common"
	"go-hexagonal/api/v1/product/request"
	"go-hexagonal/business/product"

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

	insertProductRequest := new(request.InsertProductRequest)
	if err := c.Bind(insertProductRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.InsertProduct(*insertProductRequest.ToUpsertProductSpec(int(userID)), "creator")
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
