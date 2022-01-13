package controllers

import (
	"Final_Project/config"
	"Final_Project/models/response"
	"Final_Project/models/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Userregister(c echo.Context) error {
	userRegister := users.Register{}
	c.Bind(&userRegister)
	if userRegister.Nama == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Namaoi",
			Data:    nil,
		})
	}
	var userDB users.User
	userDB.Nama = userRegister.Nama
	userDB.Alamat = userRegister.Alamat
	userDB.Umur = userRegister.Umur
	result := config.DB.Create(&userDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Namaoi",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    userDB,
	})
}
func LoginData(c echo.Context) error {
	login := users.Login{}
	c.Bind(&login)
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    login,
	})
}
func DeteilData(c echo.Context) error {
	Iduser, err := strconv.Atoi(c.Param("IdUser"))
	if err != nil {
		return c.JSON(http.StatusOK, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Tidak konvert",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    users.User{Id: Iduser},
	})
}
func GetData(c echo.Context) error {
	// nama := c.QueryParam("nama")
	// alamat := c.QueryParam("address")
	// user := User{}
	// if nama == "" {
	// 	user = User{1, "Faiq", 12, "nganjuk"}
	// } else {
	// 	user = User{1, nama, 12, alamat}
	// }
	userDB := []users.User{}
	result := config.DB.Find(&userDB)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "server",
				Data:    nil,
			})
		}
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    userDB,
	})
}
