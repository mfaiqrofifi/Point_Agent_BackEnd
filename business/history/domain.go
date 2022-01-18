package history

import (
	"Final_Project/business/product"
	"Final_Project/business/redem"
	"Final_Project/business/users"
	"context"
	"time"
)

type History struct {
	Id            int
	IdUser        int
	IdProduct     *int
	IdReward      *int
	Amount        int
	Status        string
	PoinItems     int
	Type          string
	ImageRequest  string
	NomoHp        string
	NomorRekening string
	MidtransLink  string
	Product       product.DomainProdcut
	User          users.DomainUser
	Reward        redem.DomainRedem
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Usecase interface {
	RequestProduct(ctx context.Context, idUser int, idProduct *int, amount int, img string) (History, error)
	AllowProduct(ctx context.Context, idUser int, stsatus string) (History, error)
	ViewHistoryUser(ctx context.Context, idUser int) ([]History, error)
	ViewRequestUser(ctx context.Context) ([]History, error)
	RequestRedem(ctx context.Context, idUser int, idReward *int, amount int, identyty string) (History, error)
	ViewRedem(ctx context.Context) ([]History, error)
}
type Repository interface {
	RequestProduct(ctx context.Context, idUser int, idProduct *int, amount int, img string) (History, error)
	AllowProduct(ctx context.Context, idUser int, stsatus string) (History, error)
	ViewHistoryUser(ctx context.Context, idUser int) ([]History, error)
	ViewRequestUser(ctx context.Context) ([]History, error)
	RequestRedem(ctx context.Context, idUser int, idReward *int, amount int, identity string) (History, error)
	ViewRedem(ctx context.Context) ([]History, error)
}
