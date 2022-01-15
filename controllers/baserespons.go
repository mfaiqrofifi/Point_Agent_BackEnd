package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResonse struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccesResponse(c echo.Context, data interface{}) error {
	response := BaseResonse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Berhasil"
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func NewFailResponse(c echo.Context, status int, message string) error {
	response := BaseResonse{}
	response.Meta.Status = status
	response.Meta.Message = message
	response.Data = nil
	return c.JSON(status, response)
}
