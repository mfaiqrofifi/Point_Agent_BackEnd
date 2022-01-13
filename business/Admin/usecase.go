package Admin

import (
	middlewares "Final_Project/app/middleware"
	"context"
	"errors"
	"time"
)

type UserUsecase struct {
	ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUseCase(repo Repository, time time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: time,
	}
}

func (uc *UserUsecase) Login(ctx context.Context, username string, password string) (DomainAdmin, error) {
	if username == "" {
		return DomainAdmin{}, errors.New("username empethy")
	}
	if password == "" {
		return DomainAdmin{}, errors.New("password empethy")
	}
	var err error
	loginadmin, err := uc.Repo.Login(ctx, username, password)
	if err != nil {
		return DomainAdmin{}, err
	}
	loginadmin.Token, err = uc.ConfigJWT.GenerateToken(loginadmin.Id, "admin")
	if err != nil {
		return DomainAdmin{}, err
	}
	return loginadmin, nil
}
