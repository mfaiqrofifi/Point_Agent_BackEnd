package product

import (
	"Final_Project/app/middleware"
	"Final_Project/business/product"
	"Final_Project/controllers"
	"Final_Project/controllers/product/request"
	"Final_Project/controllers/product/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllerProduct struct {
	UserUseCase product.UsecaseProduct
}

func NewUserControllerProduct(usecase product.UsecaseProduct) *UserControllerProduct {
	return &UserControllerProduct{
		UserUseCase: usecase,
	}
}

func (userController UserControllerProduct) ProductAddProduct(c echo.Context) error {
	Admin, _ := middleware.GetUser(c)
	if Admin == "admin" {
		addUser := request.RequestAdduser{}
		c.Bind(&addUser)
		ctx := c.Request().Context()
		product, error := userController.UserUseCase.AddProduct(ctx, addUser.NameProduct, addUser.Poin, addUser.Amount, addUser.Img)
		result := response.ResponseProduct{
			Id:          product.Id,
			NameProduct: product.NameProduct,
			Poin:        product.Poin,
			Amount:      product.Amount,
			Img:         product.Img,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}
		if error != nil {
			return controllers.NewFailResponse(c, http.StatusBadRequest, error.Error())
		}
		return controllers.NewSuccesResponse(c, result)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Anda Bukan Admin")
}

func (userController UserControllerProduct) ProductKind(c echo.Context) error {
	ctx := c.Request().Context()

	product, error := userController.UserUseCase.ProductKind(ctx)
	// result := response.ResponseProduct{
	// 	Id:          product.Id,
	// 	NameProduct: product.NameProduct,
	// 	Poin:        product.Poin,
	// 	Amount:      product.Amount,
	// 	Img:         product.Img,
	// 	CreatedAt:   product.CreatedAt,
	// 	UpdatedAt:   product.UpdatedAt,
	// }
	if error != nil {
		return controllers.NewFailResponse(c, http.StatusInternalServerError, error.Error())
	}
	resp := []response.Respjsn{}
	for _, user := range product {
		resp = append(resp, response.ToResponse(user))
	}
	return controllers.NewSuccesResponse(c, resp)
}

func (userController UserControllerProduct) Delete(c echo.Context) error {
	Admin, _ := middleware.GetUser(c)
	if Admin == "admin" {
		ctx := c.Request().Context()
		addUser := request.Delete{}
		c.Bind(&addUser)
		product, error := userController.UserUseCase.Delete(ctx, addUser.Id)
		if error != nil {
			return controllers.NewFailResponse(c, http.StatusInternalServerError, error.Error())
		}
		resp := []response.Respjsn{}
		for _, user := range product {
			resp = append(resp, response.ToResponse(user))
		}
		return controllers.NewSuccesResponse(c, product)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Anda Bukan Admin")
}
func (userController UserControllerProduct) Update(c echo.Context) error {
	Admin, _ := middleware.GetUser(c)
	if Admin == "admin" {
		addUser := request.Update{}
		c.Bind(&addUser)
		ctx := c.Request().Context()
		product, error := userController.UserUseCase.Update(ctx, addUser.NameProduct, addUser.Poin, addUser.Amount, addUser.Img, addUser.Id)
		result := response.ResponseProduct{
			Id:          product.Id,
			NameProduct: product.NameProduct,
			Poin:        product.Poin,
			Amount:      product.Amount,
			Img:         product.Img,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}
		if error != nil {
			return controllers.NewFailResponse(c, http.StatusBadRequest, error.Error())
		}
		return controllers.NewSuccesResponse(c, result)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Anda Bukan Admin")
}
