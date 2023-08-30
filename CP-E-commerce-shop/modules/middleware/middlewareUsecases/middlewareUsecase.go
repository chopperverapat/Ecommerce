package middlewareUsecases

import (
	"cpshop/modules/middleware/middlewareRepositories"
)

type IMiddlewareUsecase interface {
}

type middlewareUsecase struct {
	middlewareRepositories middlewareRepositories.IMiddlewareRepositories
}

func NewMiddlewareUsecase(middlewareRepositories middlewareRepositories.IMiddlewareRepositories) IMiddlewareUsecase {
	return &middlewareUsecase{
		middlewareRepositories: middlewareRepositories,
	}
}
