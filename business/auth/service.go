package auth

import (
	"fmt"
	"go-hexagonal/business"
	"go-hexagonal/business/user"
	"net/mail"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/streadway/amqp"
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

func PubliserEmail(emailAddress string) {

	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s",
		os.Getenv("GOHEXAGONAL_RABBIT_USER"),
		os.Getenv("GOHEXAGONAL_RABBIT_PASS"),
		os.Getenv("GOHEXAGONAL_RABBIT_ADDRESS"),
		os.Getenv("GOHEXAGONAL_RABBIT_PORt"))

	conn, err := amqp.Dial(connectionString)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Succes conn")

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"EmailQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"EmailQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(emailAddress),
			// Timestamp:   time.Now(),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Succesfully publish email")

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

	if user != nil {
		return "", business.ErrInvalidUsername
	}

	_, err = mail.ParseAddress(email)

	if err != nil {
		return "", business.ErrInvalidEmail
	}

	userNew := NewUser(username, email, password, firstname, lastname)
	err = s.userService.InsertUser(userNew, "system")

	if err != nil {
		return "", err
	}

	// PubliserEmail(email)

	//call notification service
	return "", err
}
