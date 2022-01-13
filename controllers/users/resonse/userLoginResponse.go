package resonse

import "time"

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
