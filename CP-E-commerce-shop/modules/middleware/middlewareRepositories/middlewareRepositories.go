package middlewareRepositories

import (
	"github.com/jmoiron/sqlx"
)

type IMiddlewareRepositories interface {
}

type middlewareRepositories struct {
	db *sqlx.DB
}

func NewMiddlewareRepositories(db *sqlx.DB) IMiddlewareRepositories {
	return &middlewareRepositories{
		db: db,
	}
}
