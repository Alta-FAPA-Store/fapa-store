package user

import "time"

//User product User that available to rent or sell
type User struct {
	ID         int
	Firstname  string
	Lastname   string
	Username   string
	Password   string
	Role       string
	Email      string
	Address    string
	Phone      string
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt time.Time
	ModifiedBy string
	Version    int
}

//NewUser create new User
func NewUser(
	id int,
	first_name string,
	last_name string,
	phone string,
	username string,
	password string,
	role string,
	email string,
	creator string,
	createdAt time.Time) User {

	return User{
		ID:         id,
		Firstname:  first_name,
		Lastname:   last_name,
		Username:   username,
		Password:   password,
		Role:       role,
		Phone:      phone,
		Email:      email,
		CreatedAt:  createdAt,
		CreatedBy:  creator,
		ModifiedAt: createdAt,
		ModifiedBy: creator,
		Version:    1,
	}
}

//ModifyUser update existing User data
func (oldData *User) ModifyUser(update UpdateUserRequest, modifiedAt time.Time, updater string) User {
	return User{
		ID:         oldData.ID,
		Firstname:  update.Firstname,
		Lastname:   update.Lastname,
		Username:   update.Username,
		Password:   oldData.Password,
		CreatedAt:  oldData.CreatedAt,
		CreatedBy:  oldData.CreatedBy,
		ModifiedAt: modifiedAt,
		ModifiedBy: updater,
		Version:    oldData.Version + 1,
	}
}
