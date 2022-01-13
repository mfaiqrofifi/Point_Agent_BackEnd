package response

import "time"

type Bankresp struct {
	Id            int       `json:"id"`
	NameType      string    `json:"nameType"`
	Img           string    `json:"img"`
	NominalReward int       `json:"nominalReward"`
	NamaBank      string    `json:"namaInstansi"`
	Poin          int       `json:"poin"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
