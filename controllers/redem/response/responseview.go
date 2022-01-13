package response

import (
	"Final_Project/business/redem"
	"time"
)

type ViewResp struct {
	Id            int       `json:"id"`
	NameType      string    `json:"nameType"`
	Img           string    `json:"img"`
	NominalReward int       `json:"nominalReward"`
	NamaOperator  string    `json:"namaOperator"`
	NamaBank      string    `json:"namaInstansi"`
	JenisEmoney   string    `json:"jenisEmoney"`
	Poin          int       `json:"poin"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func ToResponse(domain redem.DomainRedem) ViewResp {
	return ViewResp{
		Id:            domain.Id,
		NameType:      domain.NameType,
		Img:           domain.Img,
		NominalReward: domain.NominalReward,
		NamaOperator:  domain.JenisOperator,
		NamaBank:      domain.NamaBank,
		JenisEmoney:   domain.JenisEmoney,
		Poin:          domain.Poin,
		Description:   domain.Description,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}
