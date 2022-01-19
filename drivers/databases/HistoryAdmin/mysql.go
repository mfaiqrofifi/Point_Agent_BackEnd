package HistoryAdmin

import (
	"Final_Project/business/HistoreyAdmin"
	"context"

	"gorm.io/gorm"
)

type MySqlUserRepositoryProduct struct {
	Conn *gorm.DB
}

func NewMysqlUserRepositoryProduct(conn *gorm.DB) HistoreyAdmin.Repository {
	return &MySqlUserRepositoryProduct{
		Conn: conn,
	}
}

func (rep MySqlUserRepositoryProduct) HistoryAdmin(ctx context.Context) ([]HistoreyAdmin.HistoryAdmin, error) {
	var productdb []HistoryAdminDB
	result := rep.Conn.Order("id desc").Find(&productdb)
	if result.Error != nil {
		return []HistoreyAdmin.HistoryAdmin{}, result.Error
	}
	return ToListDeteil(productdb), nil
}
