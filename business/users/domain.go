package users

import (
	"time"

	"golang.org/x/net/context"
)

type DomainUser struct {
	Id        int
	Toko      string
	Email     string
	Password  string
	Poin      int
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Register(ctx context.Context, toko string, email string, password string, poin int) (DomainUser, error)
	LoginUser(ctx context.Context, email string, password string) (DomainUser, error)
	DeteilUser(ctx context.Context) ([]DomainUser, error)
	User(ctx context.Context, userId int) (DomainUser, error)
}

type Repository interface {
	Register(ctx context.Context, toko string, email string, password string, poin int) (DomainUser, error)
	LoginUser(ctx context.Context, email string) (DomainUser, error)
	DeteilUser(ctx context.Context) ([]DomainUser, error)
	User(ctx context.Context, userId int) (DomainUser, error)
}
