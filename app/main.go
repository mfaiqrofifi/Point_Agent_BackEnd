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
