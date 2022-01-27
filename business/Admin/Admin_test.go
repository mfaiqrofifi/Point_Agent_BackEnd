package Admin_test

import (
	_jwt "Final_Project/app/middleware"
	"Final_Project/business/Admin"
	_mockUser "Final_Project/business/Admin/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepo _mockUser.Repository
var userService Admin.Usecase
var user Admin.DomainAdmin
var useJWT _jwt.ConfigJWT

func setup() {
	userService = Admin.NewUserUseCase(&userRepo, time.Hour*1, useJWT)
	user = Admin.DomainAdmin{
		Id:       1,
		UserName: "admin",
		Password: "Pass",
		Token:    "token",
	}
}

func TestRegister(t *testing.T) {
	setup()
	userRepo.On("Login", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).Return(user, nil).Once()
	t.Run("Test Case 1|valid Login", func(t *testing.T) {
		usert, err := userService.Login(context.Background(), "admin", "Pass")
		usert.Token = "token"
		assert.Nil(t, err)
		assert.Equal(t, usert, user)
	})
	t.Run("Test Case 2|invalid Login", func(t *testing.T) {
		usert, err := userService.Login(context.Background(), "", "Pass")
		usert.Token = "token"
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3|invalid Login", func(t *testing.T) {
		usert, err := userService.Login(context.Background(), "admin", "")
		usert.Token = "token"
		assert.NotNil(t, err)
	})
}
