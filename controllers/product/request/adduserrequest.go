package request

type RequestAdduser struct {
	NameProduct string `json:"productName"`
	Poin        int    `json:"poin"`
	Amount      string `json:"amount"`
	Img         string `json:"img"`
}
