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

type MidtransCustomerDetails struct {
	FirstName    string
	LastName     string
	Email        string
	Phone        string
	Address      string
	TotalPayment int
}

func NewTransaction(userId int, cartId int, courier string, paymentMethod string, paymentUrl sql.NullString, totalPrice float32, status string, createdAt time.Time) Transaction {
	return Transaction{
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
