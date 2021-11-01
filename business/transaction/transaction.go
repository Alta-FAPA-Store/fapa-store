package transaction

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id            int
	UserId        int
	CartId        int
	Courier       string
	PaymentMethod string
	PaymentUrl    sql.NullString
	TotalPrice    float32
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UpdatedBy     sql.NullString
	DeletedAt     gorm.DeletedAt
}

type MidtransCreatePaymentRequest struct {
	Firstname     string
	Lastname      string
	Email         string
	Phone         string
	Address       string
	TransactionId int
	TotalPayment  int
	Items         []MidtransItemDetails
}

type MidtransItemDetails struct {
	Name     string `json:"product_name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func NewTransaction(id int, userId int, cartId int, courier string, paymentMethod string, paymentUrl sql.NullString, totalPrice float32, status string, createdAt time.Time) Transaction {
	return Transaction{
		Id:            id,
		UserId:        userId,
		CartId:        cartId,
		Courier:       courier,
		PaymentMethod: paymentMethod,
		PaymentUrl:    paymentUrl,
		TotalPrice:    totalPrice,
		Status:        status,
		CreatedAt:     createdAt,
		UpdatedAt:     createdAt,
	}
}
