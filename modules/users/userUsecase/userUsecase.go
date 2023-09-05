package userUsecase

import (
	"cpshop/config"
	"cpshop/modules/users"
	"cpshop/modules/users/userRepositories"
)

type IuserUsercase interface {
	InsertCustomer(uq *users.UserRegisterReq) (*users.UserPasssport, error)
}

type userUsecase struct {
	cfg              config.Iconfig
	userrepositories userRepositories.IusersRepository
}

func UserUsecase(cfg config.Iconfig, userrepositories userRepositories.IusersRepository) IuserUsercase {
	return &userUsecase{
		cfg:              cfg,
		userrepositories: userrepositories,
	}
}

func (uu *userUsecase) InsertCustomer(uq *users.UserRegisterReq) (*users.UserPasssport, error) {
	// Hashing a password
	if err := uq.BcrypPass(); err != nil {
		return nil, err
	}

	// Insert user
	result, err := uu.userrepositories.InsertUsers(uq, false)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("result_usecase: %v,\n", result)
	return result, nil
}
