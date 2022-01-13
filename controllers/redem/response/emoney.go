package response

import "time"

type EmoneyResp struct {
	Id            int       `json:"id"`
	NameType      string    `json:"nameType"`
	Img           string    `json:"img"`
	NominalReward int       `json:"nominalReward"`
	NamaEmoney    string    `json:"namaEmoney"`
	Poin          int       `json:"poin"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
