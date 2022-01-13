package redem

import (
	"Final_Project/business/redem"
	"time"

	"gorm.io/gorm"
)

type RedemDB struct {
	Id            int `gorm:"primaryKey"`
	NameType      string
	Img           string
	NominalReward int
	NamaBank      string
	JenisOperator string
	JenisEmoney   string
	Poin          int
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (redemDB *RedemDB) ToDomain() redem.DomainRedem {
	return redem.DomainRedem{
		Id:            redemDB.Id,
		NameType:      redemDB.NameType,
		Img:           redemDB.Img,
		NominalReward: redemDB.NominalReward,
		NamaBank:      redemDB.NamaBank,
		JenisOperator: redemDB.JenisOperator,
		JenisEmoney:   redemDB.JenisEmoney,
		Poin:          redemDB.Poin,
		Description:   redemDB.Description,
		CreatedAt:     redemDB.CreatedAt,
		UpdatedAt:     redemDB.UpdatedAt,
	}
}
func FromDomain(domain redem.DomainRedem) RedemDB {
	return RedemDB{
		Id:            domain.Id,
		NameType:      domain.NameType,
		Img:           domain.Img,
		NominalReward: domain.NominalReward,
		NamaBank:      domain.NamaBank,
		JenisOperator: domain.JenisOperator,
		JenisEmoney:   domain.JenisEmoney,
		Poin:          domain.Poin,
		Description:   domain.Description,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}
func ToListDeteil(data []RedemDB) (result []redem.DomainRedem) {
	result = []redem.DomainRedem{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}
