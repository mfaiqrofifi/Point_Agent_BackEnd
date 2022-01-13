package product

import (
	"Final_Project/business/product"
	"context"

	"gorm.io/gorm"
)

type MySqlUserRepositoryProduct struct {
	Conn *gorm.DB
}

func NewMysqlUserRepositoryProduct(conn *gorm.DB) product.RepositoryProduct {
	return &MySqlUserRepositoryProduct{
		Conn: conn,
	}
}

func (rep MySqlUserRepositoryProduct) AddProduct(ctx context.Context, nameProduct string, poin int, amount string, img string) (product.DomainProdcut, error) {
	var productdb ProductDB
	productdb.NameProduct = nameProduct
	productdb.Poin = poin
	productdb.Amount = amount
	productdb.Img = img
	result := rep.Conn.Create(&productdb)
	if result.Error != nil {
		return product.DomainProdcut{}, result.Error
	}
	return productdb.ToDomain(), nil
}

func (rep MySqlUserRepositoryProduct) ProductKind(ctx context.Context) ([]product.DomainProdcut, error) {
	var productdb []ProductDB
	result := rep.Conn.Order("id desc").Find(&productdb)
	if result.Error != nil {
		return []product.DomainProdcut{}, result.Error
	}
	return ToListDeteil(productdb), nil
}

func (rep MySqlUserRepositoryProduct) Delete(ctx context.Context, id int) ([]product.DomainProdcut, error) {
	var productdb []ProductDB
	var productdb1 ProductDB
	result := rep.Conn.Where("id = ?", id).Delete(&productdb1)
	if result.Error != nil {
		return []product.DomainProdcut{}, result.Error
	}
	result1 := rep.Conn.Order("id desc").Find(&productdb)
	if result1.Error != nil {
		return []product.DomainProdcut{}, result.Error
	}
	return ToListDeteil(productdb), nil
}

func (rep MySqlUserRepositoryProduct) Update(ctx context.Context, nameProduct string, poin int, amount string, img string, id int) (product.DomainProdcut, error) {
	var productdb ProductDB
	productdb.NameProduct = nameProduct
	productdb.Poin = poin
	productdb.Amount = amount
	productdb.Img = img
	result := rep.Conn.Model(&productdb).Where("id = ?", id).Save(&productdb)
	if result.Error != nil {
		return product.DomainProdcut{}, result.Error
	}
	return productdb.ToDomain(), nil
}
