package userhandler

import (
	"server/config"
	"server/modules/users/userUsecase"
)

type IuserHandler interface{

}

type userhandler struct {
	cfg config.Icongig
	userUsecase userUsecase.IuserUsercase
}

func Userhandler(cfg config.Icongig , userUsecase userUsecase.IuserUsercase) IuserHandler {
	return &userhandler{

	}
}