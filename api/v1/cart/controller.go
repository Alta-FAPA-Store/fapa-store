package cart

import (
	"go-hexagonal/api/common"
	"go-hexagonal/api/v1/cart/request"
	"go-hexagonal/api/v1/cart/response"
	"go-hexagonal/business/cart"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service cart.Service
}

func NewController(service cart.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) FindCartByUserId(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	userId, _ := strconv.Atoi(c.Param("user_id"))

	cart, err := controller.service.FindCartByUserId(userId)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	if cart == nil {
		return c.JSON(common.NewSuccessResponseWithoutData())
	}

	response := response.NewGetCartResponse(*cart)

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) InsertCart(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	insertCartRequest := new(request.InsertCartRequest)

	if err := c.Bind(insertCartRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.InsertCart(*insertCartRequest.ToUpsertCartSpec(insertCartRequest.UserId, insertCartRequest.ProductId))

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) DeleteCartDetails(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	deleteCartDetailsRequest := new(request.DeleteCartDetailsRequest)

	if err := c.Bind(deleteCartDetailsRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.DeleteCartDetails(*deleteCartDetailsRequest.ToUpsetDeleteCartDetailsSpec(deleteCartDetailsRequest.CartId, deleteCartDetailsRequest.ProductId))

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) UpdateQuantityCartDetails(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	updateCartDetailsResponse := new(request.UpdateCartDetailsResponse)

	if err := c.Bind(updateCartDetailsResponse); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.UpdateQuantityCartDetails(*updateCartDetailsResponse.ToUpsertUpdateCartDetailsSpec())

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
