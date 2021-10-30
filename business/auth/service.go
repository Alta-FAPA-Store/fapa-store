package auth

import (
	"go-hexagonal/business"
	"go-hexagonal/business/user"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

//=============== The implementation of those interface put below =======================
type service struct {
	userService user.Service
}

func NewUser(username, email, password, firstname, lastname string) user.InsertUserSpec {

	return user.InsertUserSpec{
		Firstname: firstname,
		Lastname:  lastname,
		Username:  username,
		Email:     email,
		Password:  password,
		Phone:     "",
	}
}

//NewService Construct user service object
func NewService(userService user.Service) Service {
	return &service{
		userService,
	}
}

//Login by given user Username and Password, return error if not exist
func (s *service) Login(username string, password string) (string, error) {
	user, err := s.userService.FindUserByUsernameAndPassword(username, password)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	claims["name"] = user.Firstname
	claims["role"] = user.Role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (s *service) Register(username, email, password, firstname, lastname string) (string, error) {
	user, err := s.userService.FindUserByUsername(username)

	if user == nil {
		return "", business.ErrInvalidUsername
	}

	userNew := NewUser(username, email, password, firstname, lastname)

	err = s.userService.InsertUser(userNew, "system")

	if err != nil {
		return "", err
	}

	//call notification service

	return "id verifikasi", err
}
