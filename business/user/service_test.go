package user_test

import (
	"go-hexagonal/business"
	"go-hexagonal/business/user"
	userMock "go-hexagonal/business/user/mocks"

	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id         = 1
	first_name = "first_name"
	last_name  = "last_name"
	phoneNum   = "0818080"
	email      = "email"
	username   = "username"
	password   = "password"
	creator    = "creator"

	modifier = "modifier"
	version  = 1
)

var (
	userService    user.Service
	userRepository userMock.Repository

	userData       user.User
	insertUserData user.InsertUserSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindUserByUsername(t *testing.T) {
	t.Run("Expect found the user", func(t *testing.T) {
		userRepository.On("FindUserByUsername", mock.AnythingOfType("string")).Return(&userData, nil).Once()

		user, err := userService.FindUserByUsername(username)

		assert.Nil(t, err)

		assert.NotNil(t, user)

		assert.Equal(t, id, user.ID)
		assert.Equal(t, first_name, user.Firstname)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)

	})

	t.Run("Expect user not found", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		user, err := userService.FindUserByUsername(username)

		assert.NotNil(t, err)

		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestInsertUserByID(t *testing.T) {
	t.Run("Expect insert user success", func(t *testing.T) {
		userRepository.On("InsertUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("string")).Return(nil).Once()

		err := userService.InsertUser(insertUserData, creator)

		assert.Nil(t, err)

	})

	t.Run("Expect insert user not found", func(t *testing.T) {
		userRepository.On("InsertUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("string")).Return(business.ErrInternalServerError).Once()

		err := userService.InsertUser(insertUserData, creator)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})
}

func TestUpdateUserByID(t *testing.T) {
	t.Run("Expect update user success", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("UpdateUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("int")).Return(nil).Once()

		err := userService.UpdateUser(username, user.UpdateUserRequest{}, modifier, version)

		assert.Nil(t, err)

	})

	t.Run("Expect update user failed", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("UpdateUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("int")).Return(business.ErrInternalServerError).Once()

		err := userService.UpdateUser(username, user.UpdateUserRequest{}, modifier, version)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})
}

func setup() {

	userData = user.NewUser(
		id,
		first_name,
		last_name,
		phoneNum,
		username,
		password,
		email,
		creator,
		time.Now(),
	)

	insertUserData = user.InsertUserSpec{
		Firstname: first_name,
		Lastname:  last_name,
		Phone:     phoneNum,
		Username:  username,
		Password:  password,
	}

	userService = user.NewService(&userRepository)
}
