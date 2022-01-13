package request

type Pulsa struct {
	NameType      string `json:"nameType"`
	Img           string `json:"img"`
	NominalReward int    `json:"nominalReward"`
	NamaOperator  string `json:"namaOperator"`
	Poin          int    `json:"poin"`
	Description   string `json:"description"`
}
