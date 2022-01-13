package routes

import (
	"Final_Project/controllers"

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	ev1 := e.Group("v1/")
	ev1.GET("users", controllers.GetData)
	ev1.GET("users/:IdUser", controllers.DeteilData)
	ev1.POST("login", controllers.LoginData)
	ev1.POST("register", controllers.Userregister)
	return e
}
