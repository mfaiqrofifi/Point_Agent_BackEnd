package users

import (
	"Final_Project/business/users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        int    `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Toko      string `gorm:"unique"`
	Password  string
	Poin      int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (user *Users) ToDomain() users.DomainUser {
	return users.DomainUser{
		Id:        user.Id,
		Email:     user.Email,
		Toko:      user.Toko,
		Password:  user.Password,
		Poin:      user.Poin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FromDomain(domain users.DomainUser) Users {
	return Users{
		Id:        domain.Id,
		Email:     domain.Email,
		Toko:      domain.Toko,
		Password:  domain.Password,
		Poin:      domain.Poin,
		UpdatedAt: domain.UpdatedAt,
		CreatedAt: domain.CreatedAt,
	}
}
func ToListDeteil(data []Users) (result []users.DomainUser) {
	result = []users.DomainUser{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}
