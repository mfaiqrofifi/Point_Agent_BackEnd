package request

type Bank struct {
	NameType      string `json:"nameType"`
	Img           string `json:"img"`
	NominalReward int    `json:"nominalReward"`
	NamaBank      string `json:"namaBank"`
	Poin          int    `json:"poin"`
	Description   string `json:"description"`
}
