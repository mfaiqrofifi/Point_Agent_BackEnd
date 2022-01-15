package resonse

import (
	"Final_Project/business/users"
	"time"
)

type UserLoginResponse struct {
	Id        int       `json:"id"`
	Toko      string    `json:"toko"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Poin      int       `json:"poin"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToResponse(domain users.DomainUser) UserLoginResponse {
	return UserLoginResponse{
		Id:        domain.Id,
		Toko:      domain.Toko,
		Email:     domain.Email,
		Password:  domain.Password,
		Poin:      domain.Poin,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
