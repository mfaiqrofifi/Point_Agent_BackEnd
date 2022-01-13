package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int            `gorm:"primaryKey" json:"id"`
	Nama      string         `json:"nama"`
	Umur      int            `json:"umur"`
	Alamat    string         `json:"alamat"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"uppdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Register struct {
	Nama   string `json:"nama"`
	Umur   int    `json:"umur"`
	Alamat string `json:"alamat"`
}
