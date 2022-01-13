package Admin

import (
	"context"
	"time"
)

type DomainAdmin struct {
	Id        int
	UserName  string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Login(ctx context.Context, username string, password string) (DomainAdmin, error)
}

type Repository interface {
	Login(ctx context.Context, username string, password string) (DomainAdmin, error)
}
