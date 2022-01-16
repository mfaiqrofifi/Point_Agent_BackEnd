package redem

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

func (uc *UserUsecase) AddRedemBank(ctx context.Context, nameType string, img string, nominalReward int, namaBank string, poin int, description string) (DomainRedem, error) {
	if nameType == "" {
		return DomainRedem{}, errors.New("type name is empethy")
	}
	if img == "" {
		return DomainRedem{}, errors.New("img is empethy")
	}
	// if nominalReward == 0 {
	// 	return DomainRedem{}, errors.New("nominal reward is empethy")
	// }
	if namaBank == "" {
		return DomainRedem{}, errors.New("nama bank is empethy")
	}
	if poin == 0 {
		return DomainRedem{}, errors.New("poin is empethy")
	}
	if description == "" {
		return DomainRedem{}, errors.New("descriptiom is empethy")
	}
	addRedem, err := uc.Repo.AddRedemBank(ctx, nameType, img, nominalReward, namaBank, poin, description)
	if err != nil {
		return DomainRedem{}, err
	}
	return addRedem, nil
}

func (uc *UserUsecase) AddRedemPulsa(ctx context.Context, nameType string, img string, nominalReward int, jenisOperator string, poin int, description string) (DomainRedem, error) {
	if nameType == "" {
		return DomainRedem{}, errors.New("type name is empethy")
	}
	if img == "" {
		return DomainRedem{}, errors.New("img is empethy")
	}
	if nominalReward == 0 {
		return DomainRedem{}, errors.New("nominal reward is empethy")
	}
	if jenisOperator == "" {
		return DomainRedem{}, errors.New("nama operator is empethy")
	}
	if poin == 0 {
		return DomainRedem{}, errors.New("poin is empethy")
	}
	if description == "" {
		return DomainRedem{}, errors.New("descriptiom is empethy")
	}
	addRedem, err := uc.Repo.AddRedemBank(ctx, nameType, img, nominalReward, jenisOperator, poin, description)
	if err != nil {
		return DomainRedem{}, err
	}
	return addRedem, nil
}

func (uc *UserUsecase) AddRedemEmoney(ctx context.Context, nameType string, img string, nominalReward int, jenisEmoney string, poin int, description string) (DomainRedem, error) {
	if nameType == "" {
		return DomainRedem{}, errors.New("type name is empethy")
	}
	if img == "" {
		return DomainRedem{}, errors.New("img is empethy")
	}
	// if nominalReward == 0 {
	// 	return DomainRedem{}, errors.New("nominal reward is empethy")
	// }
	if jenisEmoney == "" {
		return DomainRedem{}, errors.New("nama emoney is empethy")
	}
	if poin == 0 {
		return DomainRedem{}, errors.New("poin is empethy")
	}
	if description == "" {
		return DomainRedem{}, errors.New("descriptiom is empethy")
	}
	addRedem, err := uc.Repo.AddRedemBank(ctx, nameType, img, nominalReward, jenisEmoney, poin, description)
	if err != nil {
		return DomainRedem{}, err
	}
	return addRedem, nil
}
func (uc *UserUsecase) ViewRedem(ctx context.Context) ([]DomainRedem, error) {
	user, err := uc.Repo.ViewRedem(ctx)
	if err != nil {
		return []DomainRedem{}, err
	}
	return user, nil
}
func (uc *UserUsecase) DeleteRedem(ctx context.Context, id int) ([]DomainRedem, error) {
	user, err := uc.Repo.DeleteRedem(ctx, id)
	if err != nil {
		return []DomainRedem{}, err
	}
	return user, nil
}
func (uc *UserUsecase) UpdateRedem(ctx context.Context, nameType string, img string, nominalReward int, namaBank string, poin int, description string, id int) (DomainRedem, error) {
	if nameType == "" {
		return DomainRedem{}, errors.New("type name is empethy")
	}
	if img == "" {
		return DomainRedem{}, errors.New("img is empethy")
	}
	// if nominalReward == 0 {
	// 	return DomainRedem{}, errors.New("nominal reward is empethy")
	// }
	if namaBank == "" {
		return DomainRedem{}, errors.New("nama bank is empethy")
	}
	if poin == 0 {
		return DomainRedem{}, errors.New("poin is empethy")
	}
	if description == "" {
		return DomainRedem{}, errors.New("descriptiom is empethy")
	}
	addRedem, err := uc.Repo.UpdateRedem(ctx, nameType, img, nominalReward, namaBank, poin, description, id)
	if err != nil {
		return DomainRedem{}, err
	}
	return addRedem, nil
}
