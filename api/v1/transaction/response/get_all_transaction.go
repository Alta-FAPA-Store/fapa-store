package response

import (
	"go-hexagonal/api/paginator"
	"go-hexagonal/business/transaction"
)

type GetAllTransactionResponse struct {
	Meta         paginator.Meta                  `json:"meta"`
	Transactions []GetTransactionDetailsResponse `json:"transactions"`
}

func NewAllTransactionResponse(transactions []transaction.Transaction, page int, rowPerPage int) GetAllTransactionResponse {
	var lenTransactions = len(transactions)

	getAllTransactionResponse := GetAllTransactionResponse{}
	getAllTransactionResponse.Meta.BuildMeta(lenTransactions, page, rowPerPage)

	for _, value := range transactions {
		var transaction GetTransactionDetailsResponse

		transaction.Id = value.Id
		transaction.UserId = value.UserId
		transaction.CartId = value.CartId
		transaction.Courier = value.Courier
		transaction.PaymentMethod = value.PaymentMethod
		transaction.PaymentUrl = value.PaymentUrl
		transaction.TotalPrice = value.TotalPrice
		transaction.Status = value.Status
		transaction.CreatedAt = value.CreatedAt

		getAllTransactionResponse.Transactions = append(getAllTransactionResponse.Transactions, transaction)
	}

	if getAllTransactionResponse.Transactions == nil {
		getAllTransactionResponse.Transactions = []GetTransactionDetailsResponse{}
	}

	return getAllTransactionResponse
}
