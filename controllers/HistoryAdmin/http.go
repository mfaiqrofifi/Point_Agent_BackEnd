package Historyadmin

import (
	"Final_Project/app/middleware"
	"Final_Project/business/HistoreyAdmin"
	"Final_Project/controllers"
	"Final_Project/controllers/HistoryAdmin/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllerProduct struct {
	UserUseCase HistoreyAdmin.Usecase
}

func NewUserControllerProduct(usecase HistoreyAdmin.Usecase) *UserControllerProduct {
	return &UserControllerProduct{
		UserUseCase: usecase,
	}
}

func (userController UserControllerProduct) HistoreyAdmin(c echo.Context) error {
	ctx := c.Request().Context()
	Admin, _ := middleware.GetUser(c)
	if Admin == "admin" {
		product, error := userController.UserUseCase.HistoryAdmin(ctx)
		if error != nil {
			return controllers.NewFailResponse(c, http.StatusInternalServerError, error.Error())
		}
		resp := []response.AdminHist{}
		for _, user := range product {
			resp = append(resp, response.ToResponse(user))
		}
		return controllers.NewSuccesResponse(c, resp)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Anda Bukan Admin")
}
