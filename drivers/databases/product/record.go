package product

import (
	"Final_Project/business/product"
	"time"

	"gorm.io/gorm"
)

type ProductDB struct {
	Id          int `gorm:"primaryKey"`
	NameProduct string
	Poin        int
	Amount      string
	Img         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (productDB *ProductDB) ToDomain() product.DomainProdcut {
	return product.DomainProdcut{
		Id:          productDB.Id,
		NameProduct: productDB.NameProduct,
		Poin:        productDB.Poin,
		Amount:      productDB.Amount,
		Img:         productDB.Img,
		CreatedAt:   productDB.CreatedAt,
		UpdatedAt:   productDB.UpdatedAt,
	}
}

func FromDomain(domainproduct product.DomainProdcut) ProductDB {
	return ProductDB{
		Id:          domainproduct.Id,
		NameProduct: domainproduct.NameProduct,
		Poin:        domainproduct.Poin,
		Amount:      domainproduct.Amount,
		Img:         domainproduct.Img,
		CreatedAt:   domainproduct.CreatedAt,
		UpdatedAt:   domainproduct.UpdatedAt,
	}
}

func ToListDeteil(data []ProductDB) (result []product.DomainProdcut) {
	result = []product.DomainProdcut{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}
