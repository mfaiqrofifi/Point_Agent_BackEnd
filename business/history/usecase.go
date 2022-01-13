package history

import (
	"context"
	"errors"
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

func (uc *UserUsecase) RequestProduct(ctx context.Context, idUser int, idProduct *int, amount int, img string) (History, error) {
	if amount == 0 {
		return History{}, errors.New("amount is zero")
	}
	if img == "" {
		return History{}, errors.New("image is emphety")
	}
	addProduct, err := uc.Repo.RequestProduct(ctx, idUser, idProduct, amount, img)
	if err != nil {
		return History{}, err
	}
	return addProduct, nil
}
func (uc *UserUsecase) AllowProduct(ctx context.Context, idUser int, stsatus string) (History, error) {
	if idUser == 0 {
		return History{}, errors.New("idUser is zero")
	}
	if stsatus == "" {
		return History{}, errors.New("status is emphety")
	}
	allowProduct, err := uc.Repo.AllowProduct(ctx, idUser, stsatus)
	if err != nil {
		return History{}, err
	}
	return allowProduct, err
}

func (uc *UserUsecase) ViewHistoryUser(ctx context.Context, idUser int) ([]History, error) {
	user, err := uc.Repo.ViewHistoryUser(ctx, idUser)
	if err != nil {
		return []History{}, err
	}
	return user, nil
}

func (uc *UserUsecase) ViewRequestUser(ctx context.Context) ([]History, error) {
	user, err := uc.Repo.ViewRequestUser(ctx)
	if err != nil {
		return []History{}, err
	}
	return user, nil
}

func (uc *UserUsecase) RequestRedem(ctx context.Context, idUser int, idReward *int, amount int, identity string) (History, error) {
	if idUser == 0 {
		return History{}, errors.New("idUser is zero")
	}
	// if idReward == nil {
	// 	return History{}, errors.New("idreward is emphety")
	// }
	if amount == 0 {
		return History{}, errors.New("amount is zero")
	}
	allowProduct, err := uc.Repo.RequestRedem(ctx, idUser, idReward, amount, identity)
	if err != nil {
		return History{}, err
	}
	return allowProduct, err
}

func (uc *UserUsecase) ViewRedem(ctx context.Context) ([]History, error) {
	user, err := uc.Repo.ViewRedem(ctx)
	if err != nil {
		return []History{}, err
	}
	return user, nil
}
