package request

type Update struct {
	Id            int    `json:"id"`
	NameType      string `json:"nameType"`
	Img           string `json:"img"`
	NominalReward int    `json:"nominalReward"`
	NamaBank      string `json:"namaBank"`
	Poin          int    `json:"poin"`
	Description   string `json:"description"`
}
