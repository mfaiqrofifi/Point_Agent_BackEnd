package request

type UserRegisterReq struct {
	Toko     string `json:"toko"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Poin     int    `json:"poin"`
}
