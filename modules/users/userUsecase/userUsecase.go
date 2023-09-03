package userusecase

import (
	"server/config"
	"server/modules/users/userRepositories"
)

type IuserUsercase interface {
}

type userUsecase struct {
	cfg              config.Icongig
	userrepositories userRepositories.IusersRepository
}

func UserUsecase(cfg config.Icongig, userrepositories userRepositories.IusersRepository) IuserUsercase {
	return &userUsecase{
		cfg:              cfg,
		userrepositories: userrepositories,
	}
}
