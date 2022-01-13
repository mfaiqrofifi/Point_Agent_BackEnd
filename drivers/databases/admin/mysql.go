package admin

import (
	"Final_Project/business/Admin"
	"context"

	"gorm.io/gorm"
)

type MySqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) Admin.Repository {
	return &MySqlUserRepository{
		Conn: conn,
	}
}

func (rep MySqlUserRepository) Login(ctx context.Context, username string, password string) (Admin.DomainAdmin, error) {
	var admin Admins
	result := rep.Conn.First(&admin, "username = ? AND password = ?", username, password)
	if result.Error != nil {
		return Admin.DomainAdmin{}, result.Error
	}
	return admin.ToDomain(), nil
}
