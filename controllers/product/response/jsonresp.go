package response

import (
	"Final_Project/business/product"
	"time"
)

type Respjsn struct {
	Id          int       `json:"id"`
	NameProduct string    `json:"nameProduct"`
	Poin        int       `json:"poin"`
	Amount      string    `json:"amount"`
	Img         string    `json:"img"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"deletedAt"`
}

func ToResponse(domain product.DomainProdcut) Respjsn {
	return Respjsn{
		Id:          domain.Id,
		NameProduct: domain.NameProduct,
		Poin:        domain.Poin,
		Amount:      domain.Amount,
		Img:         domain.Img,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
