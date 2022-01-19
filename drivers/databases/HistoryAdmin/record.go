package HistoryAdmin

import (
	"Final_Project/business/HistoreyAdmin"
	"time"
)

type HistoryAdminDB struct {
	Id          int `gorm:"primaryKey"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (history *HistoryAdminDB) ToDomain() HistoreyAdmin.HistoryAdmin {
	return HistoreyAdmin.HistoryAdmin{
		Id:          history.Id,
		Description: history.Description,
		CreatedAt:   history.CreatedAt,
		UpdatedAt:   history.UpdatedAt,
	}
}

func FromDomain(domain HistoreyAdmin.HistoryAdmin) HistoryAdminDB {
	return HistoryAdminDB{
		Id:          domain.Id,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
func ToListDeteil(data []HistoryAdminDB) (result []HistoreyAdmin.HistoryAdmin) {
	result = []HistoreyAdmin.HistoryAdmin{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}
