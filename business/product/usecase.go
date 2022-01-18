package product

import (
	
	"context"
	"errors"
	"time"
)

type UserUsecase struct {
	Repo           RepositoryProduct
	contextTimeout time.Duration
}

func NewUserUseCase(repo RepositoryProduct, time time.Duration) UsecaseProduct {
	return &UserUsecase{
		Repo:           repo,
		contextTimeout: time,
	}
}

func (uc *UserUsecase) AddProduct(ctx context.Context, nameProduct string, poin int, amount string, img string) (DomainProdcut, error) {
	if nameProduct == "" {
		return DomainProdcut{}, errors.New("product name empethy")
	}
	if amount == "" {
		return DomainProdcut{}, errors.New("amount empethy")
	}
	if img == "" {
		return DomainProdcut{}, errors.New("product name empethy")
	}
	
	addProduct, err := uc.Repo.AddProduct(ctx, nameProduct, poin, amount, img)
	if err != nil {
		return DomainProdcut{}, err
	}
	return addProduct, nil
}

func (uc *UserUsecase) ProductKind(ctx context.Context) ([]DomainProdcut, error) {
	user, err := uc.Repo.ProductKind(ctx)
	if err != nil {
		return []DomainProdcut{}, err
	}
	return user, nil
}

func (uc *UserUsecase) Delete(ctx context.Context, id int) ([]DomainProdcut, error) {
	user, err := uc.Repo.Delete(ctx, id)
	if err != nil {
		return []DomainProdcut{}, err
	}
	return user, nil
}

func (uc *UserUsecase) Update(ctx context.Context, nameProduct string, poin int, amount string, img string, id int) (DomainProdcut, error) {
	if nameProduct == "" {
		return DomainProdcut{}, errors.New("product name empethy")
	}
	if amount == "" {
		return DomainProdcut{}, errors.New("amount empethy")
	}
	if img == "" {
		return DomainProdcut{}, errors.New("product name empethy")
	}
	addProduct, err := uc.Repo.Update(ctx, nameProduct, poin, amount, img, id)
	if err != nil {
		return DomainProdcut{}, err
	}
	return addProduct, nil
}
