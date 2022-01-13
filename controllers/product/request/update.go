package request

type Update struct {
	Id          int    `json:"id"`
	NameProduct string `json:"productName"`
	Poin        int    `json:"poin"`
	Amount      string `json:"amount"`
	Img         string `json:"img"`
}
