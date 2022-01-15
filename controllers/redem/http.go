package redem

import (
	"Final_Project/app/middleware"
	"Final_Project/business/redem"
	"Final_Project/controllers"
	"Final_Project/controllers/redem/request"
	"Final_Project/controllers/redem/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ControlleRedem struct {
	UserUseCase redem.Usecase
}

func NewUserControllerRedem(usecase redem.Usecase) *ControlleRedem {
	return &ControlleRedem{
		UserUseCase: usecase,
	}
}

func (Controller ControlleRedem) AddRedemBank(c echo.Context) error {
	Admin, _ := middleware.GetUser(c)
	if Admin == "admin" {
		addredembank := request.Bank{}
		c.Bind(&addredembank)
		ctx := c.Request().Context()
		redem, error := Controller.UserUseCase.AddRedemBank(ctx, addredembank.NameType, addredembank.Img, addredembank.NominalReward, addredembank.NamaBank, addredembank.Poin, addredembank.Description)
		if error != nil {
			return controllers.NewFailResponse(c, http.StatusInternalServerError, error.Error())
		}
		result := response.Bankresp{
			Id:            redem.Id,
			NameType:      redem.NameType,
			Img:           redem.Img,
			NominalReward: redem.NominalReward,
			NamaBank:      redem.NamaBank,
			Poin:          redem.Poin,
			Description:   redem.Description,
			CreatedAt:     redem.CreatedAt,
			UpdatedAt:     redem.UpdatedAt,
		}
		return controllers.NewSuccesResponse(c, result)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Anda Bukan Admin")
}

func (Controller ControlleRedem) AddRedemPulsa(c echo.Context) error {
	Admin, _ := middleware.GetUser(c)
	if Admin == "admin" {
		addredempulsa := request.Pulsa{}
		c.Bind(&addredempulsa)
		ctx := c.Request().Context()
		redempulsa, error := Controller.UserUseCase.AddRedemPulsa(ctx, addredempulsa.NameType, addredempulsa.Img, addredempulsa.NominalReward, addredempulsa.NamaOperator, addredempulsa.Poin, addredempulsa.Description)
		if error != nil {
			return controllers.NewFailResponse(c, http.StatusInternalServerError, error.Error())
		}
		result := response.PulsaResp{
			Id:            redempulsa.Id,
			NameType:      redempulsa.NameType,
			Img:           redempulsa.Img,
			NominalReward: redempulsa.NominalReward,
			NamaOperator:  redempulsa.JenisOperator,
			Poin:          redempulsa.Poin,
			Description:   redempulsa.Description,
			CreatedAt:     redempulsa.CreatedAt,
			UpdatedAt:     redempulsa.UpdatedAt,
		}
		return controllers.NewSuccesResponse(c, result)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Anda Bukan Admin")
}

func (Controller ControlleRedem) AddRedemEmoney(c echo.Context) error {
	Admin, _ := middleware.GetUser(c)
	if Admin == "admin" {
		addredememoney := request.Emoney{}
		c.Bind(&addredememoney)
		ctx := c.Request().Context()
		redememoney, error := Controller.UserUseCase.AddRedemEmoney(ctx, addredememoney.NameType, addredememoney.Img, addredememoney.NominalReward, addredememoney.NamaEmoney, addredememoney.Poin, addredememoney.Description)
		if error != nil {
			return controllers.NewFailResponse(c, http.StatusInternalServerError, error.Error())
		}
		result := response.PulsaResp{
			Id:            redememoney.Id,
			NameType:      redememoney.NameType,
			Img:           redememoney.Img,
			NominalReward: redememoney.NominalReward,
			NamaOperator:  redememoney.JenisOperator,
			Poin:          redememoney.Poin,
			Description:   redememoney.Description,
			CreatedAt:     redememoney.CreatedAt,
			UpdatedAt:     redememoney.UpdatedAt,
		}
		return controllers.NewSuccesResponse(c, result)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Anda Bukan Admin")
}
func (userController ControlleRedem) ViewRedem(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := userController.UserUseCase.ViewRedem(ctx)
	if err != nil {
		return controllers.NewFailResponse(c, http.StatusInternalServerError, err.Error())
	}
	resp := []response.ViewResp{}
	for _, user := range users {
		resp = append(resp, response.ToResponse(user))
	}
	return controllers.NewSuccesResponse(c, resp)
}

func (userController ControlleRedem) Delete(c echo.Context) error {
	Admin, _ := middleware.GetUser(c)
	if Admin == "admin" {
		ctx := c.Request().Context()
		addUser := request.DeleteRedem{}
		c.Bind(&addUser)
		product, error := userController.UserUseCase.DeleteRedem(ctx, addUser.Id)
		if error != nil {
			return controllers.NewFailResponse(c, http.StatusInternalServerError, error.Error())
		}
		resp := []response.ViewResp{}
		for _, user := range product {
			resp = append(resp, response.ToResponse(user))
		}
		return controllers.NewSuccesResponse(c, resp)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Anda Bukan Admin")
}

func (Controller ControlleRedem) Update(c echo.Context) error {
	Admin, _ := middleware.GetUser(c)
	if Admin == "admin" {
		addredembank := request.Update{}
		c.Bind(&addredembank)
		ctx := c.Request().Context()
		redem, error := Controller.UserUseCase.UpdateRedem(ctx, addredembank.NameType, addredembank.Img, addredembank.NominalReward, addredembank.NamaBank, addredembank.Poin, addredembank.Description, addredembank.Id)
		if error != nil {
			return controllers.NewFailResponse(c, http.StatusInternalServerError, error.Error())
		}
		result := response.Bankresp{
			Id:            redem.Id,
			NameType:      redem.NameType,
			Img:           redem.Img,
			NominalReward: redem.NominalReward,
			NamaBank:      redem.NamaBank,
			Poin:          redem.Poin,
			Description:   redem.Description,
			CreatedAt:     redem.CreatedAt,
			UpdatedAt:     redem.UpdatedAt,
		}
		return controllers.NewSuccesResponse(c, result)
	}
	return controllers.NewFailResponse(c, http.StatusUnauthorized, "Anda Bukan Admin")
}
