package request

type UserRegisterReqEdit struct {
	Id       int    `json:"id"`
	Toko     string `json:"toko"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Poin     int    `json:"poin"`
}
