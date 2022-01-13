package request

type Emoney struct {
	NameType      string `json:"nameType"`
	Img           string `json:"img"`
	NominalReward int    `json:"nominalReward"`
	NamaEmoney    string `json:"namaEmoney"`
	Poin          int    `json:"poin"`
	Description   string `json:"description"`
}
