package transaction

import (
	"go-hexagonal/api/common"
	"go-hexagonal/api/paginator"
	"go-hexagonal/api/v1/transaction/request"
	"go-hexagonal/api/v1/transaction/response"
	"go-hexagonal/business/transaction"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"strconv"
)

type Controller struct {
	service transaction.Service
}

func NewController(service transaction.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) CreateTransaction(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	createTransactionRequest := new(request.CreateTransactionRequest)

	if err := c.Bind(createTransactionRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	urlRedirect, err := controller.service.CreateTransaction(*createTransactionRequest.ToUpSertTransactionSpec())

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponse(urlRedirect))
}

func (controller *Controller) GetAllTransaction(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	userId, _ := strconv.Atoi(c.QueryParam("user_id"))
	limit := c.QueryParam("limit")
	offset := c.QueryParam("page")

	skip, page, rowPerPage := paginator.CreatePagination(offset, limit)

	transactions, err := controller.service.GetAllTransaction(userId, rowPerPage-1, skip)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewAllTransactionResponse(transactions, page, rowPerPage)

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) GetTransactionDetails(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	transactionId, _ := strconv.Atoi(c.Param("transaction_id"))

	transactionDetails, err := controller.service.GetTransactionDetails(transactionId)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewTransactionDetailsResponse(*transactionDetails)

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) UpdateTransaction(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	transactionId, _ := strconv.Atoi(c.Param("transaction_id"))

	updateTransactionRequest := new(request.UpdateTransactionRequest)

	if err := c.Bind(updateTransactionRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.UpdateTransaction(transactionId, updateTransactionRequest.Status)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) DeleteTransaction(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	transactionId, _ := strconv.Atoi(c.Param("transaction_id"))

	err := controller.service.DeleteTransaction(transactionId)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
