package users

import (
	middlewares "Final_Project/app/middleware"
	"Final_Project/helpers/encript"
	"errors"
	"fmt"
	"time"

	"golang.org/x/net/context"
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

func (uc *UserUsecase) Register(ctx context.Context, toko string, email string, password string, poin int) (DomainUser, error) {
	if email == "" {
		return DomainUser{}, errors.New("email empethy")
	}
	if password == "" {
		return DomainUser{}, errors.New("password empethy")
	}
	if toko == "" {
		return DomainUser{}, errors.New("toko empethy")
	}
	var err error
	password, err = encript.Hash(password)
	if err != nil {
		return DomainUser{}, err
	}
	registeruser, err := uc.Repo.Register(ctx, toko, email, password, poin)
	if err != nil {
		return DomainUser{}, err
	}
	fmt.Println(registeruser)
	return registeruser, nil
}
func (uc *UserUsecase) LoginUser(ctx context.Context, email string, password string) (DomainUser, error) {

	if email == "" {
		return DomainUser{}, errors.New("email empethy")
	}
	if password == "" {
		return DomainUser{}, errors.New("password empethy")
	}
	var err error
	loginuser, err := uc.Repo.LoginUser(ctx, email)
	temp := encript.ValidateHash(password, loginuser.Password)
	if !temp {
		return DomainUser{}, errors.New("password wrong")
	}
	if err != nil {
		return DomainUser{}, err
	}
	loginuser.Token, err = uc.ConfigJWT.GenerateToken(loginuser.Id, "user")
	if err != nil {
		return DomainUser{}, err
	}
	return loginuser, nil
}

func (uc *UserUsecase) DeteilUser(ctx context.Context) ([]DomainUser, error) {
	user, err := uc.Repo.DeteilUser(ctx)
	if err != nil {
		return []DomainUser{}, err
	}
	return user, nil
}

func (uc *UserUsecase) User(ctx context.Context, userId int) (DomainUser, error) {
	user, err := uc.Repo.User(ctx, userId)
	if err != nil {
		return DomainUser{}, err
	}
	return user, nil
}
