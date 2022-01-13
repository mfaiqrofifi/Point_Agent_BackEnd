package admin

import (
	"Final_Project/business/Admin"
	"time"
)

type Admins struct {
	Id        int `gorm:"primaryKey"`
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (adminDB *Admins) ToDomain() Admin.DomainAdmin {
	return Admin.DomainAdmin{
		Id:        adminDB.Id,
		UserName:  adminDB.UserName,
		Password:  adminDB.Password,
		CreatedAt: adminDB.CreatedAt,
		UpdatedAt: adminDB.UpdatedAt,
	}
}

func FromDomain(domain Admin.DomainAdmin) Admins {
	return Admins{
		Id:        domain.Id,
		UserName:  domain.UserName,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
