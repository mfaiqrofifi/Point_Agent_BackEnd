package response

import "time"

type ResponseAdmin struct {
	Id        int       `json:"id"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
