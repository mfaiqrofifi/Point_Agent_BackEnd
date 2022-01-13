package admin

import (
	"Final_Project/business/Admin"
	"Final_Project/controllers"
	"Final_Project/controllers/admin/request"
	"Final_Project/controllers/admin/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase Admin.Usecase
}

func NewUserController(userUsecase Admin.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUsecase,
	}
}
func (adminUsecase UserController) Login(c echo.Context) error {
	userLogin := request.ReqAdmin{}
	c.Bind(&userLogin)
	ctx := c.Request().Context()
	admin, err := adminUsecase.UserUseCase.Login(ctx, userLogin.UserName, userLogin.Password)
	result := response.ResponseAdmin{
		Id:       admin.Id,
		UserName: admin.UserName,
		Password: admin.Password,
		Token:    admin.Token,
	}
	if err != nil {
		return controllers.NewFailResponse(c, http.StatusInternalServerError, err.Error())
	}
	return controllers.NewSuccesResponse(c, result)
}
