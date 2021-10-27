package product_test

// func TestFindProductByID(t *testing.T) {
// 	t.Run("Expect found the user", func(t *testing.T) {
// 		userRepository.On("FindUserByUsername", mock.AnythingOfType("string")).Return(&userData, nil).Once()

// 		user, err := userService.FindUserByUsername(username)

// 		assert.Nil(t, err)

// 		assert.NotNil(t, user)

// 		assert.Equal(t, id, user.ID)
// 		assert.Equal(t, first_name, user.Firstname)
// 		assert.Equal(t, username, user.Username)
// 		assert.Equal(t, password, user.Password)

// 	})

// 	t.Run("Expect user not found", func(t *testing.T) {
// 		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

// 		user, err := userService.FindUserByUsername(username)

// 		assert.NotNil(t, err)

// 		assert.Nil(t, user)

// 		assert.Equal(t, err, business.ErrNotFound)
// 	})
// }
