package history_test

import (
	"Final_Project/business/history"
	_mockUser "Final_Project/business/history/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepo _mockUser.Repository
var userService history.Usecase
var hsitorys history.History

func setup() {
	userService = history.NewUserUseCase(&userRepo, time.Hour*1)
	hsitorys = history.History{
		Id:            1,
		IdUser:        1,
		Amount:        1,
		ImageRequest:  "picture",
		NomoHp:        "01290129",
		NomorRekening: "121287192",
	}
}

func TestRequestProduct(t *testing.T) {
	setup()
	userRepo.On("RequestProduct", mock.Anything,
		mock.AnythingOfType("int"),
		mock.Anything,
		mock.AnythingOfType("int"),
		mock.AnythingOfType("string")).Return(hsitorys, nil).Once()
	t.Run("Test Case 1|Valid RequestProduct", func(t *testing.T) {
		_, err := userService.RequestProduct(context.Background(), 1, nil, 1, "gambar")
		assert.Nil(t, err)
	})
	// t.Run("Test Case 2|Valid RequestProduct", func(t *testing.T) {
	// 	_, err := userService.RequestProduct(context.Background(), 0, nil, 1, "gambar")
	// 	assert.NotNil(t, err)
	// })
	t.Run("Test Case 3|Valid RequestProduct", func(t *testing.T) {
		_, err := userService.RequestProduct(context.Background(), 1, nil, 0, "gambar")
		assert.NotNil(t, err)
	})
	t.Run("Test Case 4|Valid RequestProduct", func(t *testing.T) {
		_, err := userService.RequestProduct(context.Background(), 1, nil, 1, "")
		assert.NotNil(t, err)
	})
}
func TestAllowProduct(t *testing.T) {
	setup()
	userRepo.On("AllowProduct", mock.Anything,
		mock.AnythingOfType("int"),
		mock.AnythingOfType("string")).Return(hsitorys, nil).Once()
	t.Run("Test Case 1|Valid AllowProduct", func(t *testing.T) {
		_, err := userService.AllowProduct(context.Background(), 1, "gambar")
		assert.Nil(t, err)
	})
	t.Run("Test Case 2|Valid AllowProduct", func(t *testing.T) {
		_, err := userService.AllowProduct(context.Background(), 0, "gambar")
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3|Valid AllowProduct", func(t *testing.T) {
		_, err := userService.AllowProduct(context.Background(), 1, "")
		assert.NotNil(t, err)
	})
}
func TestViewHistoryUser(t *testing.T) {
	setup()
	userRepo.On("ViewHistoryUser", mock.Anything,
		mock.AnythingOfType("int")).Return([]history.History{}, nil).Once()
	t.Run("Test Case 1|Valid ViewHistoryUser", func(t *testing.T) {
		_, err := userService.ViewHistoryUser(context.Background(), 1)
		assert.Nil(t, err)
	})
}

func Test(t *testing.T) {
	setup()
	userRepo.On("RequestRedem", mock.Anything,
		mock.AnythingOfType("int"),
		mock.Anything,
		mock.AnythingOfType("int"),
		mock.AnythingOfType("string")).Return(hsitorys, nil).Once()
	t.Run("Test Case 1|Valid RequestRedem", func(t *testing.T) {
		_, err := userService.RequestRedem(context.Background(), 1, nil, 1, "nomor")
		assert.Nil(t, err)
	})
	t.Run("Test Case 2|Valid RequestRedem", func(t *testing.T) {
		_, err := userService.RequestRedem(context.Background(), 0, nil, 1, "nomor")
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3|Valid RequestRedem", func(t *testing.T) {
		_, err := userService.RequestRedem(context.Background(), 1, nil, 0, "nomor")
		assert.NotNil(t, err)
	})
}
func TestViewRequest(t *testing.T) {
	setup()
	userRepo.On("ViewRequestUser", mock.Anything).Return([]history.History{}, nil).Once()
	t.Run("Test Case 1|Valid ViewRequest", func(t *testing.T) {
		_, err := userService.ViewRequestUser(context.Background())
		assert.Nil(t, err)
	})
}
func TestViewRedem(t *testing.T) {
	setup()
	userRepo.On("ViewRedem", mock.Anything).Return([]history.History{}, nil).Once()
	t.Run("Test Case 1|Valid ViewRedem", func(t *testing.T) {
		_, err := userService.ViewRedem(context.Background())
		assert.Nil(t, err)
	})
}
