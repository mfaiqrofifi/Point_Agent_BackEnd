package resonse

import "time"

type UserResponseRegist struct {
	Id        int       `json:"id"`
	Toko      string    `json:"toko"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Poin      int       `json:"poin"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
