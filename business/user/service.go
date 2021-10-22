package user

import (
	"go-hexagonal/business"
	"go-hexagonal/util/validator"
	"time"
)

//InsertUserSpec create user spec
type InsertUserSpec struct {
	Firstname string `validate:"required"`
	Lastname  string `validate:"required"`
	Email     string `validate:"required"`
	Username  string `validate:"required"`
	Password  string `validate:"required"`
	Phone     string `validate:"required"`
}

type UpdateUserRequest struct {
	Firstname string `validate:"required"`
	Lastname  string `validate:"required"`
	Email     string `validate:"required"`
	Username  string `validate:"required"`
	Phone     string `validate:"required"`
	Version   string `validate:"required"`
}

//=============== The implementation of those interface put below =======================
type service struct {
	repository Repository
}

//NewService Construct user service object
func NewService(repositoryParam Repository) Service {
	return &service{
		repositoryParam,
	}
}

//FindUserByID Get user by given ID, return nil if not exist
func (s *service) FindUserByID(id int) (*User, error) {
	return s.repository.FindUserByID(id)
} //
func (s *service) FindUserByUsername(username string) (*User, error) {
	return s.repository.FindUserByUsername(username)
}

//FindUserByUsernameAndPassword Get user by given ID, return nil if not exist
func (s *service) FindUserByUsernameAndPassword(username string, password string) (*User, error) {
	return s.repository.FindUserByUsernameAndPassword(username, password)
}

//FindAllUser Get all users , will be return empty array if no data or error occured
func (s *service) FindAllUser(skip int, rowPerPage int) ([]User, error) {

	user, err := s.repository.FindAllUser(skip, rowPerPage)
	if err != nil {
		return []User{}, err
	}

	return user, err
}

//InsertUser Create new user and store into database
func (s *service) InsertUser(insertUserSpec InsertUserSpec, createdBy string) error {
	err := validator.GetValidator().Struct(insertUserSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	user := NewUser(
		1,
		insertUserSpec.Firstname,
		insertUserSpec.Lastname,
		insertUserSpec.Phone,
		insertUserSpec.Username,
		insertUserSpec.Password,
		insertUserSpec.Email,
		createdBy,
		time.Now(),
	)

	err = s.repository.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}

//UpdateUser will update found user, if not exists will be return error
func (s *service) UpdateUser(id int, update UpdateUserRequest, modifiedBy string, currentVersion int) error {

	user, err := s.repository.FindUserByID(id)
	if err != nil {
		return err
	} else if user == nil {
		return business.ErrNotFound
	} else if user.Version != currentVersion {
		return business.ErrHasBeenModified
	}

	modifiedUser := user.ModifyUser(update, time.Now(), modifiedBy)

	return s.repository.UpdateUser(modifiedUser, currentVersion)
}
