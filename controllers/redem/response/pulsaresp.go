package response

import "time"

type PulsaResp struct {
	Id            int       `json:"id"`
	NameType      string    `json:"nameType"`
	Img           string    `json:"img"`
	NominalReward int       `json:"nominalReward"`
	NamaOperator  string    `json:"namaOperator"`
	Poin          int       `json:"poin"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
