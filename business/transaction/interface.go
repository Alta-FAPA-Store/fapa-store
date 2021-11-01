package transaction

type Service interface {
	GetAllTransaction(userId int, limit int, offset int) ([]Transaction, error)
	GetTransactionDetails(transactionId int) (*Transaction, error)
	CreateTransaction(createTransactionSpec CreateTransactionSpec) (string, error)
	UpdateTransaction(transactionId int, status string) error
	DeleteTransaction(transactionId int) error
}

type Repository interface {
	GetAllTransaction(userId int, limit int, offset int) ([]Transaction, error)
	GetTransactionDetails(transactionId int) (*Transaction, error)
	CreateTransaction(transaction Transaction) (int, error)
	UpdateTransaction(transactionId int, status string) error
	UpdatePaymentUrlWithStatusTransaction(transactionId int, paymentUrl string, status string) error
	DeleteTransaction(transactionId int) error
	GetMidtransPaymentRequest(transactionId int, createTransactionSpec CreateTransactionSpec) (MidtransCreatePaymentRequest, error)
}
