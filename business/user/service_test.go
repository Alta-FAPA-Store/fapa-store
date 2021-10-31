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
	first_name = "Fadel"
	last_name  = "Lukman"
	phoneNum   = "0818080"
	email      = "email"
	username   = "fadellh"
	password   = "123"
	role       = "user"
	creator    = "creator"
	modifier   = "modifier"
	version    = 1
)

var (
	userService    user.Service
	userRepository userMock.Repository

	userData       user.User
	userDataAll    []user.User
	insertUserData user.InsertUserSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindUserByUsernameAndPassword(t *testing.T) {
	t.Run("Expect found the user", func(t *testing.T) {
		userRepository.On("FindUserByUsernameAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&userData, nil).Once()

		user, err := userService.FindUserByUsernameAndPassword(username, password)

		assert.Nil(t, err)

		assert.NotNil(t, user)

		assert.Equal(t, id, user.ID)
		assert.Equal(t, first_name, user.Firstname)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)

	})

	t.Run("Expect user not found", func(t *testing.T) {
		userRepository.On("FindUserByUsernameAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		user, err := userService.FindUserByUsernameAndPassword(username, password)

		assert.NotNil(t, err)

		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestFindAllUser(t *testing.T) {
	t.Run("Expect found the user", func(t *testing.T) {
		userRepository.On("FindAllUser", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(&userDataAll, nil).Once()

		user, err := userService.FindAllUser(1, 10)

		assert.Nil(t, err)

		assert.NotNil(t, user)

		assert.Equal(t, id, user[0].ID)
		assert.Equal(t, first_name, user[0].Firstname)
		assert.Equal(t, username, user[0].Username)
		assert.Equal(t, password, user[0].Password)

	})

	// t.Run("Expect user not found", func(t *testing.T) {
	// 	userRepository.On("FindAllUser", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()

	// 	user, err := userService.FindAllUser(1, 10)

	// 	assert.NotNil(t, err)

	// 	assert.Nil(t, user)

	// 	assert.Equal(t, err, business.ErrNotFound)
	// })
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
		userRepository.On("FindUserByUsername", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		user, err := userService.FindUserByUsername(username)

		assert.NotNil(t, err)

		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestInsertUser(t *testing.T) {
	t.Run("Expect insert user success", func(t *testing.T) {
		userRepository.On("InsertUser", mock.AnythingOfType("user.User")).Return(nil).Once()

		err := userService.InsertUser(insertUserData, creator)

		assert.Nil(t, err)

	})

	t.Run("Expect insert user not found", func(t *testing.T) {
		userRepository.On("InsertUser", mock.AnythingOfType("user.User")).Return(business.ErrInternalServerError).Once()

		err := userService.InsertUser(insertUserData, creator)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})
}

func TestUpdateUserByUsername(t *testing.T) {
	t.Run("Expect update user success", func(t *testing.T) {
		userRepository.On("FindUserByUsername", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("UpdateUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("int")).Return(nil).Once()

		err := userService.UpdateUser(username, user.UpdateUserRequest{}, modifier, version)

		assert.Nil(t, err)

	})

	t.Run("Expect update user failed", func(t *testing.T) {
		userRepository.On("FindUserByUsername", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("UpdateUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("int")).Return(business.ErrInternalServerError).Once()

		err := userService.UpdateUser(username, user.UpdateUserRequest{}, modifier, version)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expect user not found", func(t *testing.T) {
		userRepository.On("FindUserByUsername", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		userRepository.On("UpdateUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("int")).Return(business.ErrInternalServerError).Once()

		err := userService.UpdateUser(username, user.UpdateUserRequest{}, modifier, version)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrNotFound)
	})

	t.Run("Expect user failed", func(t *testing.T) {
		userRepository.On("FindUserByUsername", mock.AnythingOfType("string")).Return(nil, business.ErrInternalServerError).Once()
		userRepository.On("UpdateUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("int")).Return(business.ErrInternalServerError).Once()

		err := userService.UpdateUser(username, user.UpdateUserRequest{}, modifier, version)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expect version not match", func(t *testing.T) {
		userRepository.On("FindUserByUsername", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("UpdateUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("int")).Return(business.ErrHasBeenModified).Once()

		err := userService.UpdateUser(username, user.UpdateUserRequest{}, modifier, version+1)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrHasBeenModified)
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
		role,
		email,
		creator,
		time.Now(),
	)

	userDataAll = append(userDataAll, userData)

	insertUserData = user.InsertUserSpec{
		Firstname: first_name,
		Lastname:  last_name,
		Phone:     phoneNum,
		Username:  username,
		Password:  password,
		Email:     email,
	}

	userService = user.NewService(&userRepository)
}
