package users

import (
	"Final_Project/business/users"
	"Final_Project/controllers"
	"Final_Project/controllers/users/request"
	"Final_Project/controllers/users/resonse"
	"fmt"
	"net/http"

	"Final_Project/app/middleware"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.Usecase
}

func NewUserController(userUsecase users.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUsecase,
	}
}

func (userController UserController) Register(c echo.Context) error {
	Admin, _ := middleware.GetUser(c)
	fmt.Println(Admin)
	if Admin == "admin" {

		userRegister := request.UserRegisterReq{}
		c.Bind(&userRegister)
		ctx := c.Request().Context()
		user, error := userController.UserUseCase.Register(ctx, userRegister.Toko, userRegister.Email, userRegister.Password, userRegister.Poin)
		result := resonse.UserResponseRegist{
			Id:       user.Id,
			Email:    user.Email,
			Password: user.Password,
			Toko:     user.Toko,
			Poin:     user.Poin,
		}

		if error != nil {
			return controllers.NewFailResponse(c, http.StatusBadRequest, error.Error())
		}
		return controllers.NewSuccesResponse(c, result)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Anda Bukan Admin")
}

func (userController UserController) LoginUser(c echo.Context) error {
	userLogin := request.LoginUser{}
	c.Bind(&userLogin)
	ctx := c.Request().Context()
	user, error := userController.UserUseCase.LoginUser(ctx, userLogin.Email, userLogin.Password)
	result := resonse.UserLoginResponse{
		Id:        user.Id,
		Email:     user.Email,
		Password:  user.Password,
		Toko:      user.Toko,
		Token:     user.Token,
		Poin:      user.Poin,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}
	if error != nil {
		return controllers.NewFailResponse(c, http.StatusBadRequest, error.Error())
	}
	return controllers.NewSuccesResponse(c, result)
}

func (userController UserController) DeteilUser(c echo.Context) error {
	Admin, _ := middleware.GetUser(c)
	if Admin == "admin" {
		ctx := c.Request().Context()
		users, error := userController.UserUseCase.DeteilUser(ctx)

		if error != nil {
			return controllers.NewFailResponse(c, http.StatusBadRequest, error.Error())
		}
		resp := []resonse.UserLoginResponse{}
		for _, user := range users {
			resp = append(resp, resonse.ToResponse(user))
		}
		return controllers.NewSuccesResponse(c, resp)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Anda Bukan Admin")
}
func (userController UserController) User(c echo.Context) error {
	User, IdUser := middleware.GetUser(c)
	if User == "user" {
		ctx := c.Request().Context()
		user, error := userController.UserUseCase.User(ctx, IdUser)
		if error != nil {
			return controllers.NewFailResponse(c, http.StatusBadRequest, error.Error())
		}
		result := resonse.UserLoginResponse{
			Id:        user.Id,
			Email:     user.Email,
			Password:  user.Password,
			Toko:      user.Toko,
			Token:     user.Token,
			Poin:      user.Poin,
			UpdatedAt: user.UpdatedAt,
			CreatedAt: user.CreatedAt,
		}
		return controllers.NewSuccesResponse(c, result)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "User Unauthorized")
}

func (userController UserController) Edit(c echo.Context) error {
	Admin, _ := middleware.GetUser(c)
	if Admin == "admin" {

		userRegister := request.UserRegisterReqEdit{}
		c.Bind(&userRegister)
		ctx := c.Request().Context()
		user, error := userController.UserUseCase.Edit(ctx, userRegister.Toko, userRegister.Email, userRegister.Password, userRegister.Poin, userRegister.Id)
		result := resonse.UserResponseRegist{
			Id:       user.Id,
			Email:    user.Email,
			Password: user.Password,
			Toko:     user.Toko,
			Poin:     user.Poin,
		}

		if error != nil {
			return controllers.NewFailResponse(c, http.StatusBadRequest, error.Error())
		}
		return controllers.NewSuccesResponse(c, result)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Anda Bukan Admin")
}
func (userController UserController) Delete(c echo.Context) error {
	Admin, _ := middleware.GetUser(c)
	if Admin == "admin" {
		ctx := c.Request().Context()
		addUser := request.UserRegisterReqDelete{}
		c.Bind(&addUser)
		product, error := userController.UserUseCase.Delete(ctx, addUser.Id)
		if error != nil {
			return controllers.NewFailResponse(c, http.StatusBadRequest, error.Error())
		}
		resp := []resonse.UserLoginResponse{}
		for _, user := range product {
			resp = append(resp, resonse.ToResponse(user))
		}
		return controllers.NewSuccesResponse(c, resp)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Anda Bukan Admin")
}
