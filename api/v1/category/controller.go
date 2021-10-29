package category

import (
	"go-hexagonal/api/common"
	"go-hexagonal/api/paginator"
	"go-hexagonal/api/v1/category/request"
	"go-hexagonal/api/v1/category/response"
	"go-hexagonal/business/category"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service category.Service
}

//NewController Construct item API controller
func NewController(service category.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) FindCategoryByID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	id, _ := strconv.Atoi(c.Param("id"))

	category, err := controller.service.FindCategoryByID(id)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetCategoryResponse(*category)

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) FindAllCategory(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	pageQueryParam := c.QueryParam("page")
	rowPerPageQueryParam := c.QueryParam("row_per_page")

	skip, page, rowPerPage := paginator.CreatePagination(pageQueryParam, rowPerPageQueryParam)

	categories, err := controller.service.FindAllCategory(skip, rowPerPage)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetAllCategoryResponse(categories, page, rowPerPage)

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) InsertCategory(c echo.Context) error {

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

	insertCategoryRequest := new(request.InsertCategoryRequest)
	if err := c.Bind(insertCategoryRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.InsertCategory(*insertCategoryRequest.ToUpsertCategorySpec(), "creator")
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) UpdateCategory(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	id, _ := strconv.Atoi(c.Param("id"))

	updateCategoryRequest := new(request.UpdateCategoryRequest)
	if err := c.Bind(updateCategoryRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.UpdateCategory(id, *updateCategoryRequest.ToUpsertCategorySpec())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) DeleteCategory(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.service.DeleteCategory(id)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
