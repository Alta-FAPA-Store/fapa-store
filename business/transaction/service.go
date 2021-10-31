package transaction

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"go-hexagonal/business"
	"go-hexagonal/util/validator"
	"net/http"
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

type ListItem struct {
	Price int
	Qty   int
	Name  string
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) CreateTransaction(createTransactionSpec CreateTransactionSpec) (string, error) {
	err := validator.GetValidator().Struct(createTransactionSpec)

	if err != nil {
		return "", business.ErrInvalidSpec
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

	transactionId, err := s.repository.CreateTransaction(transactionData)

	if err != nil {
		return "", err
	}

	var redirectUrl string
	if createTransactionSpec.PaymentMethod == "midtrans" {
		midtransCreatePaymentRequest, err := s.repository.GetMidtransPaymentRequest(transactionId, createTransactionSpec)

		if err != nil {
			return "", err
		}

		postBody, _ := json.Marshal(map[string]interface{}{
			"first_name":     midtransCreatePaymentRequest.Firstname,
			"last_name":      midtransCreatePaymentRequest.Lastname,
			"email":          midtransCreatePaymentRequest.Email,
			"phone":          midtransCreatePaymentRequest.Phone,
			"address":        midtransCreatePaymentRequest.Address,
			"transaction_id": midtransCreatePaymentRequest.TransactionId,
			"total_price":    midtransCreatePaymentRequest.TotalPayment,
			"items":          midtransCreatePaymentRequest.Items,
		})

		requestBody := bytes.NewReader(postBody)

		res, err := http.Post("http://127.0.0.1:8000/v1/payment", "application/json", requestBody)

		if err != nil {
			return "", err
		}

		var responseData map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&responseData)

		if err != nil {
			return "", nil
		}

		res.Body.Close()

		redirectUrl = responseData["data"].(string)
	}

	return redirectUrl, nil
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
