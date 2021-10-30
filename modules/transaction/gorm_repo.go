package transaction

import (
	"database/sql"
	"go-hexagonal/business/transaction"
	"time"

	"github.com/midtrans/midtrans-go"
	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

type Transaction struct {
	Id            int            `gorm:"id"`
	UserId        int            `gorm:"user_id"`
	CartId        int            `gorm:"cart_id"`
	Courier       string         `gorm:"courier"`
	PaymentMethod string         `gorm:"payment_method"`
	PaymentUrl    sql.NullString `gorm:"payment_url"`
	TotalPrice    float32        `gorm:"total_price"`
	Status        string         `gorm:"status"`
	CreatedAt     time.Time      `gorm:"created_at"`
	UpdatedAt     time.Time      `gorm:"updated_at"`
	UpdatedBy     sql.NullString `gorm:"updated_by"`
	DeletedAt     gorm.DeletedAt
}

type MidtransCustomerDetails struct {
	Firstname    string
	Lastname     string
	Email        string
	Phone        string
	Address      string
	TotalPayment int
}

type MidtransItemDetails struct {
	Name     string
	Price    int
	Quantity int
}

//NewGormDBRepository Generate Gorm DB transaction repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func NewTransactionData(transaction transaction.Transaction) *Transaction {
	return &Transaction{
		transaction.Id,
		transaction.UserId,
		transaction.CartId,
		transaction.Courier,
		transaction.PaymentMethod,
		transaction.PaymentUrl,
		transaction.TotalPrice,
		transaction.Status,
		transaction.CreatedAt,
		transaction.UpdatedAt,
		transaction.UpdatedBy,
		transaction.DeletedAt,
	}
}

func (col *Transaction) ToTransactionDetails() transaction.Transaction {
	var transactionDetails transaction.Transaction

	transactionDetails.Id = col.Id
	transactionDetails.UserId = col.UserId
	transactionDetails.CartId = col.CartId
	transactionDetails.Courier = col.Courier
	transactionDetails.PaymentMethod = col.PaymentMethod
	transactionDetails.PaymentUrl = col.PaymentUrl
	transactionDetails.TotalPrice = col.TotalPrice
	transactionDetails.Status = col.Status
	transactionDetails.CreatedAt = col.CreatedAt
	transactionDetails.UpdatedAt = col.UpdatedAt
	transactionDetails.UpdatedBy = col.UpdatedBy
	transactionDetails.DeletedAt = col.DeletedAt

	return transactionDetails
}

func (col *MidtransCustomerDetails) ToMidtransCustomerDetails(totalPrice int) transaction.MidtransCustomerDetails {
	var customerRequest transaction.MidtransCustomerDetails

	customerRequest.FirstName = col.Firstname
	customerRequest.LastName = col.Lastname
	customerRequest.Email = col.Email
	customerRequest.Phone = col.Phone
	customerRequest.Address = col.Address
	customerRequest.TotalPayment = totalPrice

	return customerRequest
}

func (repo *GormRepository) CreateTransaction(transaction transaction.Transaction) (int, error) {
	transactionData := NewTransactionData(transaction)

	err := repo.DB.Create(transactionData).Error

	if err != nil {
		return 0, err
	}

	err = repo.DB.Table("carts").Where("id = ?", transactionData.CartId).Update("is_checkout", true).Error

	if err != nil {
		return 0, err
	}

	return transactionData.Id, nil
}

func (repo *GormRepository) GetMidtransCustomerDetails(createTransactionSpec transaction.CreateTransactionSpec) (transaction.MidtransCustomerDetails, []midtrans.ItemDetails, error) {
	var customerRequest MidtransCustomerDetails
	var itemDetails []MidtransItemDetails

	err := repo.DB.Table("user_tables").Select("firstname, lastname, email, phone, address").Where("id = ?", createTransactionSpec.UserId).Find(&customerRequest).Error

	if err != nil {
		return transaction.MidtransCustomerDetails{}, nil, err
	}

	repo.DB.Table("cart_details").Select("product_tables.name, product_tables.price, cart_details.quantity").Joins("JOIN product_tables ON cart_details.product_id = product_tables.id").Where("cart_id = ?", createTransactionSpec.CartId).Scan(&itemDetails)

	var dataItem []midtrans.ItemDetails
	var totalPrice int = 0
	for _, value := range itemDetails {
		totalPrice += value.Price * value.Quantity
		dataItem = append(dataItem, midtrans.ItemDetails{Name: value.Name, Price: int64(value.Price), Qty: int32(value.Quantity)})
	}

	dataCustomer := customerRequest.ToMidtransCustomerDetails(totalPrice)

	return dataCustomer, dataItem, nil
}

func (repo *GormRepository) GetAllTransaction(userId int, limit int, offset int) ([]transaction.Transaction, error) {
	var transactions []Transaction

	if userId != 0 {
		err := repo.DB.Limit(limit).Offset(offset).Where("user_id = ?", userId).Find(&transactions).Error

		if err != nil {
			return nil, err
		}
	} else {
		err := repo.DB.Limit(limit).Offset(offset).Find(&transactions).Error

		if err != nil {
			return nil, err
		}
	}

	var results []transaction.Transaction

	for _, value := range transactions {
		results = append(results, value.ToTransactionDetails())
	}

	return results, nil
}

func (repo *GormRepository) GetTransactionDetails(transactionId int) (*transaction.Transaction, error) {
	var transactionData Transaction

	err := repo.DB.Where("id = ?", transactionId).First(&transactionData).Error

	if err != nil {
		return nil, err
	}

	transactionDetails := transactionData.ToTransactionDetails()

	return &transactionDetails, nil
}

func (repo *GormRepository) UpdateTransaction(transactionId int, status string) error {
	err := repo.DB.Model(&Transaction{}).Where("id = ?", transactionId).Update("status", status).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) UpdatePaymentUrlTransaction(transactionId int, paymentUrl string) error {
	err := repo.DB.Model(&Transaction{}).Where("id = ?", transactionId).Update("payment_url", paymentUrl).Error

	if err != nil {
		return err
	}

	return nil
}
func (repo *GormRepository) DeleteTransaction(transactionId int) error {
	err := repo.DB.Where("id = ?", transactionId).Delete(&Transaction{}).Error

	if err != nil {
		return err
	}

	return nil
}
