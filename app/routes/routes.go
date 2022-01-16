package routes

import (
	_middleware "Final_Project/app/middleware"
	"Final_Project/controllers/admin"
	"Final_Project/controllers/history"
	"Final_Project/controllers/product"
	"Final_Project/controllers/redem"
	"Final_Project/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// config := middleware.JWTConfig{
// 	Claims:     &jwtCustomClaims{},
// 	SigningKey: []byte("secret"),
// }
type Controllerlist struct {
	ConfigJWT             _middleware.ConfigJWT
	ControllerHistory     history.UserController
	ControllerAdmin       admin.UserController
	RedemController       redem.ControlleRedem
	UserController        users.UserController
	UserProductController product.UserControllerProduct
}

func (cl *Controllerlist) RoutesRegister(e *echo.Echo) {
	config := middleware.JWT([]byte(cl.ConfigJWT.SecretJWT))
	e.POST("admin/register", cl.UserController.Register, config)
	e.POST("user/login", cl.UserController.LoginUser)
	e.POST("admin/addProduct", cl.UserProductController.ProductAddProduct, config)
	e.POST("admin/addRedem", cl.RedemController.AddRedemBank, config)
	e.POST("admin/addPulsa", cl.RedemController.AddRedemPulsa, config)
	e.POST("admin/addEmoney", cl.RedemController.AddRedemEmoney, config)
	e.POST("admin/login", cl.ControllerAdmin.Login)
	e.POST("user/request", cl.ControllerHistory.RequestProduct, config)
	e.POST("admin/allowRequest", cl.ControllerHistory.AllowProduct, config)
	e.POST("user/userRedem", cl.ControllerHistory.RequestRedem, config)
	e.POST("admin/DeleteProduct", cl.UserProductController.Delete, config)
	e.POST("admin/updateProduct", cl.UserProductController.Update, config)
	e.POST("admin/deleteRedem", cl.RedemController.Delete, config)
	e.POST("admin/updateRedem", cl.RedemController.Update, config)
	e.GET("user/history", cl.ControllerHistory.ViewHistoryUser, config)
	e.GET("viewProduct", cl.UserProductController.ProductKind)
	e.GET("admin/viewUser", cl.UserController.DeteilUser, config)
	e.GET("viewRedem", cl.RedemController.ViewRedem)
	e.GET("admin/ViewRequest", cl.ControllerHistory.ViewRequestUser, config)
	e.GET("admin/viewRedemUser", cl.ControllerHistory.ViewRedem, config)
	e.GET("user/Profile", cl.UserController.User, config)

}
