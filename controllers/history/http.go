package history

import (
	"Final_Project/app/middleware"
	"Final_Project/business/history"
	"Final_Project/controllers"
	"Final_Project/controllers/history/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase history.Usecase
}

func NewUserController(userUsecase history.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUsecase,
	}
}

func (userController UserController) RequestProduct(c echo.Context) error {
	User, UserId := middleware.GetUser(c)
	if User == "user" {
		requestProduct := request.HistoryReq{}
		c.Bind(&requestProduct)
		ctx := c.Request().Context()
		history, err := userController.UserUseCase.RequestProduct(ctx, UserId, requestProduct.IdProduct, requestProduct.Amount, requestProduct.Img)
		if err != nil {
			return controllers.NewFailResponse(c, http.StatusInternalServerError, err.Error())
		}
		return controllers.NewSuccesResponse(c, history)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Login Dahulu")
}

func (userController UserController) AllowProduct(c echo.Context) error {
	User, _ := middleware.GetUser(c)
	if User == "admin" {
		requestProduct := request.Allow{}
		c.Bind(&requestProduct)
		ctx := c.Request().Context()
		history, err := userController.UserUseCase.AllowProduct(ctx, requestProduct.Id, requestProduct.Status)
		if err != nil {
			return controllers.NewFailResponse(c, http.StatusInternalServerError, err.Error())
		}
		return controllers.NewSuccesResponse(c, history)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Login Dahulu")
}

func (userController UserController) ViewHistoryUser(c echo.Context) error {
	User, UserId := middleware.GetUser(c)
	if User == "user" {
		ctx := c.Request().Context()
		histori, err := userController.UserUseCase.ViewHistoryUser(ctx, UserId)
		if err != nil {
			return controllers.NewFailResponse(c, http.StatusInternalServerError, err.Error())
		}
		return controllers.NewSuccesResponse(c, histori)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Login Dahulu")
}

func (userController UserController) ViewRequestUser(c echo.Context) error {
	User, _ := middleware.GetUser(c)
	if User == "admin" {
		ctx := c.Request().Context()
		histori, err := userController.UserUseCase.ViewRequestUser(ctx)
		if err != nil {
			return controllers.NewFailResponse(c, http.StatusInternalServerError, err.Error())
		}
		return controllers.NewSuccesResponse(c, histori)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Login Dahulu")
}

func (userController UserController) RequestRedem(c echo.Context) error {
	User, UserId := middleware.GetUser(c)
	if User == "user" {
		requestRedem := request.ProductReq{}
		c.Bind(&requestRedem)
		ctx := c.Request().Context()
		history, err := userController.UserUseCase.RequestRedem(ctx, UserId, requestRedem.IdReward, requestRedem.Amount, requestRedem.Identity)
		if err != nil {
			return controllers.NewFailResponse(c, http.StatusInternalServerError, err.Error())
		}
		return controllers.NewSuccesResponse(c, history)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Login Dahulu")
}

func (userController UserController) ViewRedem(c echo.Context) error {
	User, _ := middleware.GetUser(c)
	if User == "admin" {
		ctx := c.Request().Context()
		histori, err := userController.UserUseCase.ViewRedem(ctx)
		if err != nil {
			return controllers.NewFailResponse(c, http.StatusInternalServerError, err.Error())
		}
		return controllers.NewSuccesResponse(c, histori)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Login Dahulu")
}
