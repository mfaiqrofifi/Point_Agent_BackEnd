package response

import (
	"Final_Project/business/HistoreyAdmin"
	"time"
)

type AdminHist struct {
	Id          int       `josn:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ToResponse(domain HistoreyAdmin.HistoryAdmin) AdminHist {
	return AdminHist{
		Id:          domain.Id,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
