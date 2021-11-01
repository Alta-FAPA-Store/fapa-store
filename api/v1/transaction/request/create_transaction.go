package request

import "go-hexagonal/business/transaction"

type CreateTransactionRequest struct {
	UserId        int     `json:"user_id"`
	CartId        int     `json:"cart_id"`
	Courier       string  `json:"courier"`
	PaymentMethod string  `json:"payment_method"`
	TotalPrice    float32 `json:"total_price"`
}

func (req *CreateTransactionRequest) ToUpSertTransactionSpec() *transaction.CreateTransactionSpec {
	var createTransactionSpec transaction.CreateTransactionSpec

	createTransactionSpec.UserId = req.UserId
	createTransactionSpec.CartId = req.CartId
	createTransactionSpec.Courier = req.Courier
	createTransactionSpec.PaymentMethod = req.PaymentMethod
	createTransactionSpec.TotalPrice = req.TotalPrice

	return &createTransactionSpec
}
