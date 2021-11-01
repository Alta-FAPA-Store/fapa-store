package transaction_test

import (
	"database/sql"
	"go-hexagonal/business"
	"go-hexagonal/business/transaction"
	transactionMock "go-hexagonal/business/transaction/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id             = 1
	user_id        = 1
	cart_id        = 1
	courier        = "jne"
	payment_method = "midtrans"
	total_price    = 100000
	status         = "created"
	limit          = 10
	offset         = 0
)

var (
	transactionService    transaction.Service
	transactionRepository transactionMock.Repository

	transactionData    []transaction.Transaction
	transactionDetails transaction.Transaction

	createTransaction   transaction.CreateTransactionSpec
	midtransRequest     transaction.MidtransCreatePaymentRequest
	midtransItemDetails transaction.MidtransItemDetails
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetAllTransaction(t *testing.T) {
	t.Run("Expect found all transactions", func(t *testing.T) {
		transactionRepository.On("GetAllTransaction", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(&transactionData, nil).Once()

		transactions, err := transactionService.GetAllTransaction(0, limit, offset)

		assert.Nil(t, err)
		assert.NotNil(t, transactions)

		assert.Equal(t, user_id, transactions[0].UserId)
		assert.Equal(t, courier, transactions[0].Courier)
	})

	t.Run("Expect found user transactions", func(t *testing.T) {
		transactionRepository.On("GetAllTransaction", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(&transactionData, nil).Once()

		transactions, err := transactionService.GetAllTransaction(user_id, limit, offset)

		assert.Nil(t, err)
		assert.NotNil(t, transactions)

		assert.Equal(t, user_id, transactions[0].UserId)
		assert.Equal(t, courier, transactions[0].Courier)
	})
}

func TestGetTransactionDetails(t *testing.T) {
	t.Run("Expect found transactions", func(t *testing.T) {
		transactionRepository.On("GetTransactionDetails", mock.AnythingOfType("int")).Return(&transactionDetails, nil).Once()

		transaction, err := transactionService.GetTransactionDetails(id)

		assert.Nil(t, err)
		assert.NotNil(t, transaction)

		assert.Equal(t, user_id, transaction.UserId)
		assert.Equal(t, courier, transaction.Courier)
	})

	t.Run("Expect transaction details not found", func(t *testing.T) {
		transactionRepository.On("GetTransactionDetails", mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()

		transaction, err := transactionService.GetTransactionDetails(69)

		assert.NotNil(t, err)
		assert.Nil(t, transaction)
	})
}

func TestCreateTransaction(t *testing.T) {
	t.Run("Expect create transaction success", func(t *testing.T) {
		transactionRepository.On("GetMidtransPaymentRequest", mock.AnythingOfType("int"), mock.AnythingOfType("transaction.CreateTransactionSpec")).Return(midtransRequest, nil).Once()
		transactionRepository.On("CreateTransaction", mock.AnythingOfType("transaction.Transaction")).Return(nil).Once()

		_, err := transactionService.CreateTransaction(createTransaction)

		assert.Nil(t, err)
	})
}

func setup() {
	createTransaction = transaction.CreateTransactionSpec{
		UserId:        user_id,
		CartId:        cart_id,
		Courier:       courier,
		PaymentMethod: payment_method,
		TotalPrice:    total_price,
		Status:        status,
	}

	// midtransItemDetails = transaction.MidtransItemDetails{
	// 	Name:     "Testing",
	// 	Price:    100000,
	// 	Quantity: 1,
	// }

	// midtransRequest = transaction.MidtransCreatePaymentRequest{
	// 	Firstname:     "Rayga",
	// 	Lastname:      "Kertia",
	// 	Email:         "raygakertia1@gmail.com",
	// 	Phone:         "12341234",
	// 	Address:       "Bermesiah street",
	// 	TransactionId: id,
	// 	TotalPayment:  total_price,
	// 	Items: []transaction.MidtransItemDetails{
	// 		midtransItemDetails,
	// 	},
	// }

	transactionDetails = transaction.NewTransaction(
		id,
		user_id,
		cart_id,
		courier,
		payment_method,
		sql.NullString{},
		total_price,
		status,
		time.Now(),
	)

	transactionData = []transaction.Transaction{
		transactionDetails,
	}

	transactionService = transaction.NewService(&transactionRepository)
}
