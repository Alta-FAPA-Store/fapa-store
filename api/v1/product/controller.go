package product

import (
	"go-hexagonal/api/common"
	"go-hexagonal/api/paginator"
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

func (controller *Controller) FindAllProduct(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	pageQueryParam := c.QueryParam("page")
	rowPerPageQueryParam := c.QueryParam("row_per_page")
	nameQueryParam := c.QueryParam("name")
	categoryQueryParam := c.QueryParam("category")

	skip, page, rowPerPage := paginator.CreatePagination(pageQueryParam, rowPerPageQueryParam)

	products, err := controller.service.FindAllProduct(skip, rowPerPage, categoryQueryParam, nameQueryParam)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetAllProductResponse(products, page, rowPerPage)

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) InsertProduct(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	userRole, ok := claims["role"]

	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	if userRole != "admin" {
		return c.JSON(common.NewUngrantResponse())
	}

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

func (controller *Controller) UpdateProduct(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	id, _ := strconv.Atoi(c.Param("id"))

	updateProductRequest := new(request.UpdateProductRequest)
	if err := c.Bind(updateProductRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.UpdateProduct(id, *updateProductRequest.ToUpsertProductSpec())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) DeleteProduct(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.service.DeleteProduct(id)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
