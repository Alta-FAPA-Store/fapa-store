package transaction

import (
	"database/sql"
	"errors"
	"go-hexagonal/business"
	"go-hexagonal/util/validator"
	"strconv"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
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

var s snap.Client

func initializeSnapClient() {
	s.New("SB-Mid-server-0qpp_T4NqLWf8ifdV4kJoKhl", midtrans.Sandbox)
}

func createTransaction(orderId string, customerDetails MidtransCustomerDetails, itemDetails []midtrans.ItemDetails) (string, error) {
	// Send request to Midtrans Snap API
	resp, err := s.CreateTransaction(GenerateSnapReq(orderId, customerDetails, itemDetails))
	if err != nil {
		return "", errors.New(err.GetMessage())
	}

	return resp.RedirectURL, nil
}

func GenerateSnapReq(orderId string, customerDetails MidtransCustomerDetails, itemDetails []midtrans.ItemDetails) *snap.Request {

	// Initiate Customer address
	custAddress := &midtrans.CustomerAddress{
		FName:   customerDetails.FirstName,
		LName:   customerDetails.LastName,
		Phone:   customerDetails.Phone,
		Address: customerDetails.Address,
	}

	// Initiate Snap Request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: int64(customerDetails.TotalPayment),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    customerDetails.FirstName,
			LName:    customerDetails.LastName,
			Email:    customerDetails.Email,
			Phone:    customerDetails.Phone,
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items:           &itemDetails,
	}

	return snapReq
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

	transactionId, err := s.repository.CreateTransaction(transactionData)

	if err != nil {
		return err
	}

	if createTransactionSpec.PaymentMethod == "midtrans" {
		customerDetails, itemDetails, err := s.repository.GetMidtransCustomerDetails(createTransactionSpec)

		if err != nil {
			return err
		}

		initializeSnapClient()
		paymentUrl, err := createTransaction(strconv.Itoa(transactionId), customerDetails, itemDetails)

		if err != nil {
			return err
		}

		err = s.repository.UpdatePaymentUrlTransaction(transactionId, paymentUrl)

		if err != nil {
			return err
		}
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
