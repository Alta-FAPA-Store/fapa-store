package transaction

import "github.com/midtrans/midtrans-go"

type Service interface {
	GetAllTransaction(userId int, limit int, offset int) ([]Transaction, error)
	GetTransactionDetails(transactionId int) (*Transaction, error)
	CreateTransaction(createTransactionSpec CreateTransactionSpec) error
	UpdateTransaction(transactionId int, status string) error
	DeleteTransaction(transactionId int) error
}

type Repository interface {
	GetAllTransaction(userId int, limit int, offset int) ([]Transaction, error)
	GetTransactionDetails(transactionId int) (*Transaction, error)
	CreateTransaction(transaction Transaction) (int, error)
	UpdateTransaction(transactionId int, status string) error
	UpdatePaymentUrlTransaction(transactionId int, paymentUrl string) error
	DeleteTransaction(transactionId int) error
	GetMidtransCustomerDetails(createTransactionSpec CreateTransactionSpec) (MidtransCustomerDetails, []midtrans.ItemDetails, error)
}
