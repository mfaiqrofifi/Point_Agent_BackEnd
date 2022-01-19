package HistoreyAdmin

import (
	"context"
	"time"
)

type UserUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUseCase(repo Repository, time time.Duration) Usecase {
	return &UserUsecase{
		Repo:           repo,
		contextTimeout: time,
	}
}

func (uc *UserUsecase) HistoryAdmin(ctx context.Context) ([]HistoryAdmin, error) {
	user, err := uc.Repo.HistoryAdmin(ctx)
	if err != nil {
		return []HistoryAdmin{}, err
	}
	return user, nil
}
