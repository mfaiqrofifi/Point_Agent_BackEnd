package redem

import (
	"Final_Project/business/redem"
	"context"

	"gorm.io/gorm"
)

type MySqlRedemRepository struct {
	Conn *gorm.DB
}

func NewMysqlRedemRepository(conn *gorm.DB) redem.Repository {
	return &MySqlRedemRepository{
		Conn: conn,
	}
}

func (rep MySqlRedemRepository) AddRedemBank(ctx context.Context, nameType string, img string, nominalReward int, namaBank string, poin int, description string) (redem.DomainRedem, error) {
	var redemdb RedemDB
	redemdb.NameType = nameType
	redemdb.Img = img
	redemdb.NominalReward = nominalReward
	redemdb.NamaBank = namaBank
	redemdb.Poin = poin
	redemdb.Description = description
	result := rep.Conn.Create(&redemdb)
	if result.Error != nil {
		return redem.DomainRedem{}, result.Error
	}
	return redemdb.ToDomain(), nil
}

func (rep MySqlRedemRepository) AddRedemPulsa(ctx context.Context, nameType string, img string, nominalReward int, namaOperator string, poin int, description string) (redem.DomainRedem, error) {
	var redemdb RedemDB
	redemdb.NameType = nameType
	redemdb.Img = img
	redemdb.NominalReward = nominalReward
	redemdb.NamaBank = namaOperator
	redemdb.Poin = poin
	redemdb.Description = description
	result := rep.Conn.Create(&redemdb)
	if result.Error != nil {
		return redem.DomainRedem{}, result.Error
	}
	return redemdb.ToDomain(), nil
}
func (rep MySqlRedemRepository) AddRedemEmoney(ctx context.Context, nameType string, img string, nominalReward int, jenisEmoney string, poin int, description string) (redem.DomainRedem, error) {
	var redemdb RedemDB
	redemdb.NameType = nameType
	redemdb.Img = img
	redemdb.NominalReward = nominalReward
	redemdb.JenisEmoney = jenisEmoney
	redemdb.Poin = poin
	redemdb.Description = description
	result := rep.Conn.Create(&redemdb)
	if result.Error != nil {
		return redem.DomainRedem{}, result.Error
	}
	return redemdb.ToDomain(), nil
}
func (rep MySqlRedemRepository) ViewRedem(ctx context.Context) ([]redem.DomainRedem, error) {
	var redemdb []RedemDB
	result := rep.Conn.Order("id desc").Find(&redemdb)
	if result.Error != nil {
		return []redem.DomainRedem{}, result.Error
	}
	return ToListDeteil(redemdb), nil
}

func (rep MySqlRedemRepository) DeleteRedem(ctx context.Context, id int) ([]redem.DomainRedem, error) {
	var productdb []RedemDB
	var productdb1 RedemDB
	result := rep.Conn.Where("id = ?", id).Delete(&productdb1)
	if result.Error != nil {
		return []redem.DomainRedem{}, result.Error
	}
	result1 := rep.Conn.Order("id desc").Find(&productdb)
	if result1.Error != nil {
		return []redem.DomainRedem{}, result.Error
	}
	return ToListDeteil(productdb), nil

}

func (rep MySqlRedemRepository) UpdateRedem(ctx context.Context, nameType string, img string, nominalReward int, namaBank string, poin int, description string, id int) (redem.DomainRedem, error) {
	var redemdb RedemDB
	redemdb.NameType = nameType
	redemdb.Img = img
	redemdb.NominalReward = nominalReward
	redemdb.NamaBank = namaBank
	redemdb.Poin = poin
	redemdb.Description = description
	result := rep.Conn.Model(&redemdb).Where("id = ?", id).Save(&redemdb)
	if result.Error != nil {
		return redem.DomainRedem{}, result.Error
	}
	return redemdb.ToDomain(), nil
}
