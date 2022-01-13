package users

import (
	"Final_Project/business/users"
	"context"

	"gorm.io/gorm"
)

type MySqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MySqlUserRepository{
		Conn: conn,
	}
}

func (rep MySqlUserRepository) Register(ctx context.Context, toko string, email string, password string, poin int) (users.DomainUser, error) {
	var user Users
	user.Toko = toko
	user.Email = email
	user.Password = password
	user.Poin = poin
	result := rep.Conn.Create(&user)
	if result.Error != nil {
		return users.DomainUser{}, result.Error
	}
	return user.ToDomain(), nil
}

func (rep MySqlUserRepository) LoginUser(ctx context.Context, email string) (users.DomainUser, error) {
	var user Users
	result := rep.Conn.First(&user, "email = ?", email)
	if result.Error != nil {
		return users.DomainUser{}, result.Error
	}
	return user.ToDomain(), nil
}

func (rep MySqlUserRepository) DeteilUser(ctx context.Context) ([]users.DomainUser, error) {
	var usersdb []Users
	result := rep.Conn.Order("id desc").Find(&usersdb)
	if result.Error != nil {
		return []users.DomainUser{}, result.Error
	}
	return ToListDeteil(usersdb), nil
}

func (rep MySqlUserRepository) User(ctx context.Context, userId int) (users.DomainUser, error) {
	var user Users
	result := rep.Conn.First(&user, "id=?", userId)
	if result.Error != nil {
		return users.DomainUser{}, result.Error
	}
	return user.ToDomain(), nil
}
