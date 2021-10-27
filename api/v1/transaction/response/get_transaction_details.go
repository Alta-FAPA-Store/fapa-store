package response

import (
	"database/sql"
	"go-hexagonal/business/transaction"
	"time"
)

type GetTransactionDetailsResponse struct {
	Id            int            `json:"id"`
	UserId        int            `json:"user_id"`
	CartId        int            `json:"cart_id"`
	Courier       string         `json:"courier"`
	PaymentMethod string         `json:"payment_method"`
	PaymentUrl    sql.NullString `json:"payment_url"`
	TotalPrice    float32        `json:"total_price"`
	Status        string         `json:"status"`
	CreatedAt     time.Time      `json:"created_at"`
}

func NewTransactionDetailsResponse(transactionDetails transaction.Transaction) *GetTransactionDetailsResponse {
	var transactionDetailsResponse GetTransactionDetailsResponse

	transactionDetailsResponse.Id = transactionDetails.Id
	transactionDetailsResponse.UserId = transactionDetails.UserId
	transactionDetailsResponse.CartId = transactionDetails.CartId
	transactionDetailsResponse.Courier = transactionDetails.Courier
	transactionDetailsResponse.PaymentMethod = transactionDetails.PaymentMethod
	transactionDetailsResponse.PaymentUrl = transactionDetails.PaymentUrl
	transactionDetailsResponse.TotalPrice = transactionDetails.TotalPrice
	transactionDetailsResponse.Status = transactionDetails.Status
	transactionDetailsResponse.CreatedAt = transactionDetails.CreatedAt

	return &transactionDetailsResponse
}
