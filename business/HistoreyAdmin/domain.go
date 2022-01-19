package HistoreyAdmin

import (
	"context"
	"time"
)

type HistoryAdmin struct {
	Id          int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	HistoryAdmin(ctx context.Context) ([]HistoryAdmin, error)
}

type Repository interface {
	HistoryAdmin(ctx context.Context) ([]HistoryAdmin, error)
}
