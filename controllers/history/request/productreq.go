package request

type ProductReq struct {
	IdReward *int   `json:"idReward"`
	Amount   int    `json:"amount"`
	Identity string `json:"nomor"`
}
