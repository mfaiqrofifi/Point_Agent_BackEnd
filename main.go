package main

import (
	"Final_Project/config"
	"Final_Project/routes"
)

// type User struct {
// 	Id        int            `gorm:"primaryKey" json:"id"`
// 	Nama      string         `json:"nama"`
// 	Umur      int            `json:"umur"`
// 	Alamat    string         `json:"alamat"`
// 	CreatedAt time.Time      `json:"createdAt"`
// 	UpdatedAt time.Time      `json:"uppdatedAt"`
// 	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
// }

// type BaseResponse struct {
// 	Code    int
// 	Message string
// 	Data    interface{}
// }

// type Login struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }
// type Register struct {
// 	Nama   string `json:"nama"`
// 	Umur   int    `json:"umur"`
// 	Alamat string `json:"alamat"`
// }

// var DB *gorm.DB

// func InitDB() {
// 	dsn := "root:yuleyek@tcp(127.0.0.1:3306)/point?charset=utf8mb4&parseTime=True&loc=Local"
// 	var err error
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("DB Failed Connect")
// 	}
// 	Migration()
// }
// func Migration() {
// 	DB.AutoMigrate(&User{})
// }
func main() {
	config.InitDB()
	e := routes.NewRoute()
	e.Start(":8000")
	// e := echo.New()
	// ev1 := e.Group("v1/")
	// ev1.GET("users", GetData)
	// ev1.GET("users/:IdUser", DeteilData)
	// ev1.POST("login", LoginData)
	// ev1.POST("register", Userregister)
	// e.Start(":8000")
}

// func Userregister(c echo.Context) error {
// 	userRegister := Register{}
// 	c.Bind(&userRegister)
// 	if userRegister.Nama == "" {
// 		return c.JSON(http.StatusBadRequest, BaseResponse{
// 			Code:    http.StatusBadRequest,
// 			Message: "Namaoi",
// 			Data:    nil,
// 		})
// 	}
// 	var userDB User
// 	userDB.Nama = userRegister.Nama
// 	userDB.Alamat = userRegister.Alamat
// 	userDB.Umur = userRegister.Umur
// 	result := DB.Create(&userDB)
// 	if result.Error != nil {
// 		return c.JSON(http.StatusInternalServerError, BaseResponse{
// 			Code:    http.StatusInternalServerError,
// 			Message: "Namaoi",
// 			Data:    nil,
// 		})
// 	}
// 	return c.JSON(http.StatusOK, BaseResponse{
// 		Code:    http.StatusOK,
// 		Message: "Berhasil",
// 		Data:    userDB,
// 	})
// }
// func LoginData(c echo.Context) error {
// 	login := Login{}
// 	c.Bind(&login)
// 	return c.JSON(http.StatusOK, BaseResponse{
// 		Code:    http.StatusOK,
// 		Message: "Berhasil",
// 		Data:    login,
// 	})
// }
// func DeteilData(c echo.Context) error {
// 	Iduser, err := strconv.Atoi(c.Param("IdUser"))
// 	if err != nil {
// 		return c.JSON(http.StatusOK, BaseResponse{
// 			Code:    http.StatusInternalServerError,
// 			Message: "Tidak konvert",
// 			Data:    nil,
// 		})
// 	}
// 	return c.JSON(http.StatusOK, BaseResponse{
// 		Code:    http.StatusOK,
// 		Message: "Berhasil",
// 		Data:    User{Id: Iduser},
// 	})
// }
// func GetData(c echo.Context) error {
// 	// nama := c.QueryParam("nama")
// 	// alamat := c.QueryParam("address")
// 	// user := User{}
// 	// if nama == "" {
// 	// 	user = User{1, "Faiq", 12, "nganjuk"}
// 	// } else {
// 	// 	user = User{1, nama, 12, alamat}
// 	// }
// 	userDB := []User{}
// 	result := DB.Find(&userDB)
// 	if result.Error != nil {
// 		if result.Error != gorm.ErrRecordNotFound {
// 			return c.JSON(http.StatusInternalServerError, BaseResponse{
// 				Code:    http.StatusInternalServerError,
// 				Message: "server",
// 				Data:    nil,
// 			})
// 		}
// 	}
// 	return c.JSON(http.StatusOK, BaseResponse{
// 		Code:    http.StatusOK,
// 		Message: "Berhasil",
// 		Data:    userDB,
// 	})
// }
