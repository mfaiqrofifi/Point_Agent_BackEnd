package users_test

import (
	_jwt "Final_Project/app/middleware"
	"Final_Project/business/users"
	_mockUser "Final_Project/business/users/mocks"
	"Final_Project/helpers/encript"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepo _mockUser.Repository
var userService users.Usecase
var user users.DomainUser
var useJWT _jwt.ConfigJWT

func setup() {
	Hash, _ := encript.Hash("pass")
	userService = users.NewUserUseCase(&userRepo, time.Hour*1, useJWT)
	user = users.DomainUser{
		Id:       1,
		Toko:     "store",
		Email:    "user@gmail.com",
		Password: Hash,
		Poin:     10,
		Token:    "token",
	}
}

func TestRegister(t *testing.T) {
	setup()
	userRepo.On("Register", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int")).Return(user, nil).Once()
	t.Run("Test Case 1|valid Register", func(t *testing.T) {
		_, usert := userService.Register(context.Background(), "store", "user@gmail.com", "pass", 10)
		assert.Nil(t, usert)
	})
	t.Run("Test Case 2|invalid Register", func(t *testing.T) {
		_, usert := userService.Register(context.Background(), "", "user@gmail.com", "pass", 10)
		assert.NotNil(t, usert)
	})
	t.Run("Test Case 3|invalid Register", func(t *testing.T) {
		_, usert := userService.Register(context.Background(), "store", "", "pass", 10)
		assert.NotNil(t, usert)
	})
	t.Run("Test Case 4|invalid Register", func(t *testing.T) {
		_, usert := userService.Register(context.Background(), "store", "user@gmail.com", "", 10)
		assert.NotNil(t, usert)
	})
}

func TestLogin(t *testing.T) {
	setup()
	userRepo.On("LoginUser", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).Return(user, nil).Once()
	t.Run("Test Case 1|valid Login", func(t *testing.T) {
		usert, err := userService.LoginUser(context.Background(), "user@gmail.com", "pass")
		usert.Token = "token"
		assert.Nil(t, err)
		assert.Equal(t, usert, user)
	})
	// t.Run("Test Case 2|invalid Login", func(t *testing.T) {
	// 	usert, err := userService.LoginUser(context.Background(), "user@gmail.com", "hello")
	// 	usert.Token = "token"
	// 	assert.NotNil(t, err)

	// })
	// t.Run("Test Case 3|invalid Login", func(t *testing.T) {
	// 	usert, err := userService.LoginUser(context.Background(), "user@gmail1.com", "pass")
	// 	usert.Token = "token"
	// 	assert.NotNil(t, err)
	// })
	t.Run("Test Case 4|invalid Login", func(t *testing.T) {
		_, err := userService.LoginUser(context.Background(), "", "pass")
		assert.NotNil(t, err)
	})
	t.Run("Test Case 5|invalid Login", func(t *testing.T) {
		_, err := userService.LoginUser(context.Background(), "user&gmail.com", "")
		assert.NotNil(t, err)
	})
}
func TestEdit(t *testing.T) {
	setup()
	userRepo.On("Edit", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("int")).Return(user, nil).Once()
	t.Run("Test Case 1|valid Edit", func(t *testing.T) {
		_, usert := userService.Edit(context.Background(), "store", "user@gmail.com", "pass", 10, 1)
		assert.Nil(t, usert)
	})
	t.Run("Test Case 2|invalid Edit", func(t *testing.T) {
		_, usert := userService.Edit(context.Background(), "", "user@gmail.com", "pass", 10, 1)
		assert.NotNil(t, usert)
	})
	t.Run("Test Case 3|invalid Edit", func(t *testing.T) {
		_, usert := userService.Edit(context.Background(), "store", "", "pass", 10, 1)
		assert.NotNil(t, usert)
	})
	t.Run("Test Case 4|invalid Edit", func(t *testing.T) {
		_, usert := userService.Edit(context.Background(), "store", "user@gmail.com", "", 10, 1)
		assert.NotNil(t, usert)
	})
}

func TestDelete(t *testing.T) {
	setup()
	userRepo.On("Delete", mock.Anything,
		mock.AnythingOfType("int")).Return([]users.DomainUser{}, nil).Once()
	t.Run("Test Case 1|valid Delete", func(t *testing.T) {
		_, usert := userService.Delete(context.Background(), 1)
		assert.Nil(t, usert)
	})
}

func TestUser(t *testing.T) {
	setup()
	userRepo.On("User", mock.Anything,
		mock.AnythingOfType("int")).Return(user, nil).Once()
	t.Run("Test Case 1|valid User", func(t *testing.T) {
		_, usert := userService.User(context.Background(), 1)
		assert.Nil(t, usert)
	})
}

func TestDeteilUser(t *testing.T) {
	setup()
	userRepo.On("DeteilUser", mock.Anything).Return([]users.DomainUser{}, nil).Once()
	t.Run("Test Case 1|valid DeteilUser", func(t *testing.T) {
		_, usert := userService.DeteilUser(context.Background())
		assert.Nil(t, usert)
	})
}
