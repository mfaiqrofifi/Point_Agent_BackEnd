package product

import (
	"context"
	"time"
)

type DomainProdcut struct {
	Id          int
	NameProduct string
	Poin        int
	Amount      string
	Img         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UsecaseProduct interface {
	AddProduct(ctx context.Context, nameProduct string, poin int, amount string, img string) (DomainProdcut, error)
	ProductKind(ctx context.Context) ([]DomainProdcut, error)
	Delete(ctx context.Context, id int) ([]DomainProdcut, error)
	Update(ctx context.Context, nameProduct string, poin int, amount string, img string, id int) (DomainProdcut, error)
}

type RepositoryProduct interface {
	AddProduct(ctx context.Context, nameProduct string, poin int, amount string, img string) (DomainProdcut, error)
	ProductKind(ctx context.Context) ([]DomainProdcut, error)
	Delete(ctx context.Context, id int) ([]DomainProdcut, error)
	Update(ctx context.Context, nameProduct string, poin int, amount string, img string, id int) (DomainProdcut, error)
}
