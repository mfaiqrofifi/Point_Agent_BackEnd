package main

import (
	_middleware "Final_Project/app/middleware"
	"Final_Project/app/routes"
	_productUsecase "Final_Project/business/product"
	_redemUsecase "Final_Project/business/redem"
	_userUsecase "Final_Project/business/users"
	_productControll "Final_Project/controllers/product"
	_redemControll "Final_Project/controllers/redem"
	_useController "Final_Project/controllers/users"
	_mySQL "Final_Project/drivers/databases/mysql"
	_productDB "Final_Project/drivers/databases/product"
	_repositoryProduct "Final_Project/drivers/databases/product"
	_redemDB "Final_Project/drivers/databases/redem"
	_repositoriesRedem "Final_Project/drivers/databases/redem"
	_repository "Final_Project/drivers/databases/users"
	_userDB "Final_Project/drivers/databases/users"
	"net/http"

	_adminusecase "Final_Project/business/Admin"
	_admincontroll "Final_Project/controllers/admin"
	_adminDB "Final_Project/drivers/databases/admin"
	_adminRepostory "Final_Project/drivers/databases/admin"

	_historiUsecase "Final_Project/business/history"
	_historiController "Final_Project/controllers/history"
	_historiDB "Final_Project/drivers/databases/history"
	_historiRepository "Final_Project/drivers/databases/history"
	"log"
	"time"

	// "Final_Project/mvp/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}
func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_adminDB.Admins{})
	db.AutoMigrate(&_userDB.Users{})
	db.AutoMigrate(&_productDB.ProductDB{})
	db.AutoMigrate(&_redemDB.RedemDB{})
	db.AutoMigrate(&_historiDB.HistoryDB{})
}
func main() {
	config := _mySQL.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	Conn := config.InitialDB()
	dbMigrate(Conn)
	e := echo.New()
	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}
	// http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "https://www.google.com")
	// 	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

	// 	if r.Method == "OPTIONS" {
	// 		w.Write([]byte("allowed"))
	// 		return
	// 	}

	// 	w.Write([]byte("hello"))
	// })
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	// type CORSConfig struct {
	// 	// Skipper defines a function to skip middleware.
	// 	Skipper Skipper

	// 	// AllowOrigin defines a list of origins that may access the resource.
	// 	// Optional. Default value []string{"*"}.
	// 	AllowOrigins []string `yaml:"allow_origins"`

	// 	// AllowOriginFunc is a custom function to validate the origin. It takes the
	// 	// origin as an argument and returns true if allowed or false otherwise. If
	// 	// an error is returned, it is returned by the handler. If this option is
	// 	// set, AllowOrigins is ignored.
	// 	// Optional.
	// 	AllowOriginFunc func(origin string) (bool, error) `yaml:"allow_origin_func"`

	// 	// AllowMethods defines a list methods allowed when accessing the resource.
	// 	// This is used in response to a preflight request.
	// 	// Optional. Default value DefaultCORSConfig.AllowMethods.
	// 	AllowMethods []string `yaml:"allow_methods"`

	// 	// AllowHeaders defines a list of request headers that can be used when
	// 	// making the actual request. This is in response to a preflight request.
	// 	// Optional. Default value []string{}.
	// 	AllowHeaders []string `yaml:"allow_headers"`

	// 	// AllowCredentials indicates whether or not the response to the request
	// 	// can be exposed when the credentials flag is true. When used as part of
	// 	// a response to a preflight request, this indicates whether or not the
	// 	// actual request can be made using credentials.
	// 	// Optional. Default value false.
	// 	AllowCredentials bool `yaml:"allow_credentials"`

	// 	// ExposeHeaders defines a whitelist headers that clients are allowed to
	// 	// access.
	// 	// Optional. Default value []string{}.
	// 	ExposeHeaders []string `yaml:"expose_headers"`

	// 	// MaxAge indicates how long (in seconds) the results of a preflight request
	// 	// can be cached.
	// 	// Optional. Default value 0.
	// 	MaxAge int `yaml:"max_age"`
	//   }
	//   DefaultCORSConfig = CORSConfig{
	// 	Skipper:      DefaultSkipper,
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	//   }
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	myRepositoryRedem := _repositoriesRedem.NewMysqlRedemRepository(Conn)
	redemUsecase := _redemUsecase.NewUserUseCase(myRepositoryRedem, timeoutContext)
	redemController := _redemControll.NewUserControllerRedem(redemUsecase)

	myRepositoryProduct := _repositoryProduct.NewMysqlUserRepositoryProduct(Conn)
	userUseCaseProduct := _productUsecase.NewUserUseCase(myRepositoryProduct, timeoutContext)
	ProductController := _productControll.NewUserControllerProduct(userUseCaseProduct)

	myRepository := _repository.NewMysqlUserRepository(Conn)
	userUsecase := _userUsecase.NewUserUseCase(myRepository, timeoutContext, configJWT)
	userController := _useController.NewUserController(userUsecase)

	myRepositoryAdmin := _adminRepostory.NewMysqlUserRepository(Conn)
	adminUsecase := _adminusecase.NewUserUseCase(myRepositoryAdmin, timeoutContext, configJWT)
	adminController := _admincontroll.NewUserController(adminUsecase)

	myRepositoryHistori := _historiRepository.NewMysqlUserRepository(Conn)
	historyUsecase := _historiUsecase.NewUserUseCase(myRepositoryHistori, timeoutContext)
	historyController := _historiController.NewUserController(historyUsecase)
	routesInit := routes.Controllerlist{
		ConfigJWT:             configJWT,
		ControllerHistory:     *historyController,
		ControllerAdmin:       *adminController,
		RedemController:       *redemController,
		UserController:        *userController,
		UserProductController: *ProductController,
	}
	routesInit.RoutesRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
