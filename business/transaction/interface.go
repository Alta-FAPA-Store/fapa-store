package transaction

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
	CreateTransaction(transaction Transaction) error
	UpdateTransaction(transactionId int, status string) error
	DeleteTransaction(transactionId int) error
}
