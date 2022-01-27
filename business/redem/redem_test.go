package redem_test

import (
	"Final_Project/business/redem"
	_mockUser "Final_Project/business/redem/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepo _mockUser.Repository
var userService redem.Usecase
var redems redem.DomainRedem

func setup() {
	userService = redem.NewUserUseCase(&userRepo, time.Hour*1)
	redems = redem.DomainRedem{
		Id:            1,
		NameType:      "Cahs out",
		Img:           "string",
		NominalReward: 1,
		NamaBank:      "BRI",
		JenisOperator: "Telkomsel",
		JenisEmoney:   "Gopay",
		Poin:          20,
		Description:   "redem",
	}
}

func TestAddRedem(t *testing.T) {
	setup()
	userRepo.On("AddRedemBank", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("string")).Return(redems, nil).Once()
	t.Run("test Case 1|Valid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemBank(context.Background(), "cahs out", "image", 1, "BRI", 2, "penukaran")
		assert.Nil(t, err)
	})
	t.Run("test Case 2|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemBank(context.Background(), "", "image", 1, "BRI", 2, "penukaran")
		assert.NotNil(t, err)
	})
	t.Run("test Case 3|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemBank(context.Background(), "pulsa", "", 1, "BRI", 2, "penukaran")
		assert.NotNil(t, err)
	})
	t.Run("test Case 4|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemBank(context.Background(), "pulsa", "image", 0, "BRI", 2, "penukaran")
		assert.NotNil(t, err)
	})
	t.Run("test Case 5|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemBank(context.Background(), "pulsa", "image", 1, "", 2, "penukaran")
		assert.NotNil(t, err)
	})
	t.Run("test Case 6|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemBank(context.Background(), "pulsa", "image", 1, "BRI", 0, "penukaran")
		assert.NotNil(t, err)
	})
	t.Run("test Case 7|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemBank(context.Background(), "pulsa", "image", 1, "BRI", 2, "")
		assert.NotNil(t, err)
	})
}

func TestAddRedemPulsa(t *testing.T) {
	setup()
	userRepo.On("AddRedemBank", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("string")).Return(redems, nil).Once()
	t.Run("test Case 1|Valid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemPulsa(context.Background(), "cahs out", "image", 1, "BRI", 2, "penukaran")
		assert.Nil(t, err)
	})
	t.Run("test Case 2|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemPulsa(context.Background(), "", "image", 1, "BRI", 2, "penukaran")
		assert.NotNil(t, err)
	})
	t.Run("test Case 3|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemPulsa(context.Background(), "pulsa", "", 1, "BRI", 2, "penukaran")
		assert.NotNil(t, err)
	})
	t.Run("test Case 4|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemPulsa(context.Background(), "pulsa", "image", 0, "BRI", 2, "penukaran")
		assert.NotNil(t, err)
	})
	t.Run("test Case 5|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemPulsa(context.Background(), "pulsa", "image", 1, "", 2, "penukaran")
		assert.NotNil(t, err)
	})
	t.Run("test Case 6|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemPulsa(context.Background(), "pulsa", "image", 1, "BRI", 0, "penukaran")
		assert.NotNil(t, err)
	})
	t.Run("test Case 7|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemPulsa(context.Background(), "pulsa", "image", 1, "BRI", 2, "")
		assert.NotNil(t, err)
	})
}
func TestAddRedemEmoney(t *testing.T) {
	setup()
	userRepo.On("AddRedemBank", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("string")).Return(redems, nil).Once()
	t.Run("test Case 1|Valid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemEmoney(context.Background(), "cahs out", "image", 1, "BRI", 2, "penukaran")
		assert.Nil(t, err)
	})
	t.Run("test Case 2|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemEmoney(context.Background(), "", "image", 1, "BRI", 2, "penukaran")
		assert.NotNil(t, err)
	})
	t.Run("test Case 3|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemEmoney(context.Background(), "pulsa", "", 1, "BRI", 2, "penukaran")
		assert.NotNil(t, err)
	})
	// t.Run("test Case 4|InValid AddRedem", func(t *testing.T) {
	// 	_, err := userService.AddRedemEmoney(context.Background(), "pulsa", "image", 0, "BRI", 2, "penukaran")
	// 	assert.NotNil(t, err)
	// })
	t.Run("test Case 5|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemEmoney(context.Background(), "pulsa", "image", 1, "", 2, "penukaran")
		assert.NotNil(t, err)
	})
	t.Run("test Case 6|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemEmoney(context.Background(), "pulsa", "image", 1, "BRI", 0, "penukaran")
		assert.NotNil(t, err)
	})
	t.Run("test Case 7|InValid AddRedem", func(t *testing.T) {
		_, err := userService.AddRedemEmoney(context.Background(), "pulsa", "image", 1, "BRI", 2, "")
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	setup()
	userRepo.On("DeleteRedem", mock.Anything, mock.AnythingOfType("int")).Return([]redem.DomainRedem{}, nil)
	t.Run("test Case 1|Valid TestDelete", func(t *testing.T) {
		_, err := userService.DeleteRedem(context.Background(), 1)
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	setup()
	userRepo.On("UpdateRedem", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int")).Return(redems, nil).Once()
	t.Run("test Case 1|Valid AddRedem", func(t *testing.T) {
		_, err := userService.UpdateRedem(context.Background(), "cahs out", "image", 1, "BRI", 2, "penukaran", 1)
		assert.Nil(t, err)
	})
	t.Run("test Case 2|InValid AddRedem", func(t *testing.T) {
		_, err := userService.UpdateRedem(context.Background(), "", "image", 1, "BRI", 2, "penukaran", 1)
		assert.NotNil(t, err)
	})
	t.Run("test Case 3|InValid AddRedem", func(t *testing.T) {
		_, err := userService.UpdateRedem(context.Background(), "pulsa", "", 1, "BRI", 2, "penukaran", 1)
		assert.NotNil(t, err)
	})
	// t.Run("test Case 4|InValid AddRedem", func(t *testing.T) {
	// 	_, err := userService.UpdateRedem(context.Background(), "pulsa", "image", 0, "BRI", 2, "penukaran", 1)
	// 	assert.NotNil(t, err)
	// })
	t.Run("test Case 5|InValid AddRedem", func(t *testing.T) {
		_, err := userService.UpdateRedem(context.Background(), "pulsa", "image", 1, "", 2, "penukaran", 1)
		assert.NotNil(t, err)
	})
	t.Run("test Case 6|InValid AddRedem", func(t *testing.T) {
		_, err := userService.UpdateRedem(context.Background(), "pulsa", "image", 1, "BRI", 0, "penukaran", 1)
		assert.NotNil(t, err)
	})
	t.Run("test Case 7|InValid AddRedem", func(t *testing.T) {
		_, err := userService.UpdateRedem(context.Background(), "pulsa", "image", 1, "BRI", 2, "", 1)
		assert.NotNil(t, err)
	})
}
