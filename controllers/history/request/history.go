package request

type HistoryReq struct {
	IdProduct *int   `json:"idProduct"`
	Amount    int    `json:"amount"`
	Img       string `json:"img"`
}
