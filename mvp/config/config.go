package config

import (
	"Final_Project/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:yuleyek@tcp(127.0.0.1:3306)/point?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB Failed Connect")
	}
	Migration()
}
func Migration() {
	DB.AutoMigrate(&users.User{})
}
