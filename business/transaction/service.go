package transaction

import (
	"database/sql"
	"go-hexagonal/business"
	"go-hexagonal/util/validator"
	"time"
)

type CreateTransactionSpec struct {
	UserId        int     `validate:"required"`
	CartId        int     `validate:"required"`
	Courier       string  `validate:"required"`
	PaymentMethod string  `validate:"required"`
	TotalPrice    float32 `validate:"required"`
	Status        string  `validate:"required"`
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) CreateTransaction(createTransactionSpec CreateTransactionSpec) error {
	err := validator.GetValidator().Struct(createTransactionSpec)

	if err != nil {
		return business.ErrInvalidSpec
	}

	transactionData := NewTransaction(
		createTransactionSpec.UserId,
		createTransactionSpec.CartId,
		createTransactionSpec.Courier,
		createTransactionSpec.PaymentMethod,
		sql.NullString{},
		createTransactionSpec.TotalPrice,
		createTransactionSpec.Status,
		time.Now(),
	)

	err = s.repository.CreateTransaction(transactionData)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAllTransaction(userId int, limit int, offset int) ([]Transaction, error) {
	transactions, err := s.repository.GetAllTransaction(userId, limit, offset)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *service) GetTransactionDetails(transactionId int) (*Transaction, error) {
	transaction, err := s.repository.GetTransactionDetails(transactionId)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *service) UpdateTransaction(transactionId int, status string) error {
	err := s.repository.UpdateTransaction(transactionId, status)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteTransaction(transactionId int) error {
	err := s.repository.DeleteTransaction(transactionId)

	if err != nil {
		return err
	}

	return nil
}
