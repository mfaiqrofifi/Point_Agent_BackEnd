package users_test

import (
	_jwt "Final_Project/app/middleware"
	"Final_Project/business/users"
	_mockUser "Final_Project/business/users/mocks"
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
	userService = users.NewUserUseCase(&userRepo, time.Hour*1, useJWT)
	user = users.DomainUser{
		Id:       1,
		Toko:     "store",
		Email:    "user@gmail.com",
		Password: "pass",
		Poin:     10,
		Token:    "token",
	}
}

func TestRegister(t *testing.T) {
	userRepo.On("Register", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int")).Return(user, nil).Once()
	t.Run("Test Case 1|valid Register", func(t *testing.T) {
		user, _ := userService.Register(context.Background(), "toko", "user@gmail.com", "pass", 10)
		assert.Nil(t, user)
	})

}

func TestLogin(t *testing.T) {
	userRepo.On("Login", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).Return(user, nil).Once()
	t.Run("Test Case 1|valid Login", func(t *testing.T) {
		usert, err := userService.LoginUser(context.Background(), "user@gmail.com", "pass")
		assert.Nil(t, err)
		assert.Equal(t, usert.Email, user.Email)
	})

}
