package product_test

import (
	"Final_Project/business/product"
	_mockUser "Final_Project/business/product/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userService product.UsecaseProduct
var products product.DomainProdcut
var userRepo _mockUser.RepositoryProduct

func setup() {
	userService = product.NewUserUseCase(&userRepo, time.Hour*1)
	products = product.DomainProdcut{
		Id:          1,
		NameProduct: "toko",
		Poin:        2,
		Amount:      "2 pcs",
		Img:         "gambar",
	}
}

func TestProduct(t *testing.T) {
	setup()
	userRepo.On("AddProduct", mock.Anything, mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).Return(products, nil).Once()
	t.Run("Test Case 1|Valid AddProduct", func(t *testing.T) {
		_, err := userService.AddProduct(context.Background(), "toko", 2, "2 pcs", "gambar")
		assert.Nil(t, err)
	})
	t.Run("Test Case 2|Invalid nameproduct", func(t *testing.T) {
		_, err := userService.AddProduct(context.Background(), "", 2, "2 pcs", "gambar")
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3|Invalid amount", func(t *testing.T) {
		_, err := userService.AddProduct(context.Background(), "toko", 2, "", "gambar")
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3|Invalid nameproduct", func(t *testing.T) {
		_, err := userService.AddProduct(context.Background(), "toko", 2, "2 pcs", "")
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	setup()
	userRepo.On("Update", mock.Anything, mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int")).Return(products, nil).Once()
	t.Run("Test Case 1|Valid AddProduct", func(t *testing.T) {
		_, err := userService.Update(context.Background(), "toko", 2, "2 pcs", "gambar", 1)
		assert.Nil(t, err)
	})
	t.Run("Test Case 2|Invalid nameproduct", func(t *testing.T) {
		_, err := userService.Update(context.Background(), "", 2, "2 pcs", "gambar", 1)
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3|Invalid amount", func(t *testing.T) {
		_, err := userService.Update(context.Background(), "toko", 2, "", "gambar", 1)
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3|Invalid nameproduct", func(t *testing.T) {
		_, err := userService.Update(context.Background(), "toko", 2, "2 pcs", "", 1)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	setup()
	userRepo.On("Delete",
		mock.Anything,
		mock.AnythingOfType("int")).Return([]product.DomainProdcut{}, nil)
	t.Run("Test Case 1|Valid Delete", func(t *testing.T) {
		_, err := userService.Delete(context.Background(), 1)
		assert.Nil(t, err)
	})
}

func TestKindProduct(t *testing.T) {
	setup()
	userRepo.On("ProductKind",
		mock.Anything).Return([]product.DomainProdcut{}, nil)
	t.Run("Test Case 1|Valid ProductKind", func(t *testing.T) {
		_, err := userService.ProductKind(context.Background())
		assert.Nil(t, err)
	})
}
