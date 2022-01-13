package redem

import (
	"context"
	"time"
)

type DomainRedem struct {
	Id            int
	NameType      string
	Img           string
	NominalReward int
	NamaBank      string
	JenisOperator string
	JenisEmoney   string
	Poin          int
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Usecase interface {
	AddRedemBank(ctx context.Context, nameType string, img string, nominalReward int, namaBank string, poin int, description string) (DomainRedem, error)
	AddRedemPulsa(ctx context.Context, nameType string, img string, nominalReward int, jenisOperator string, poin int, description string) (DomainRedem, error)
	AddRedemEmoney(ctx context.Context, nameType string, img string, nominalReward int, jenisEmoney string, poin int, description string) (DomainRedem, error)
	DeleteRedem(ctx context.Context, id int) ([]DomainRedem, error)
	UpdateRedem(ctx context.Context, nameType string, img string, nominalReward int, namaBank string, poin int, description string, id int) (DomainRedem, error)
	ViewRedem(ctx context.Context) ([]DomainRedem, error)
}
type Repository interface {
	AddRedemBank(ctx context.Context, nameType string, img string, nominalReward int, namaBank string, poin int, description string) (DomainRedem, error)
	AddRedemPulsa(ctx context.Context, nameType string, img string, nominalReward int, jenisOperator string, poin int, description string) (DomainRedem, error)
	AddRedemEmoney(ctx context.Context, nameType string, img string, nominalReward int, jenisEmoney string, poin int, description string) (DomainRedem, error)
	DeleteRedem(ctx context.Context, id int) ([]DomainRedem, error)
	UpdateRedem(ctx context.Context, nameType string, img string, nominalReward int, namaBank string, poin int, description string, id int) (DomainRedem, error)
	ViewRedem(ctx context.Context) ([]DomainRedem, error)
}
