package userrepositories

import "database/sql"

type IusersRepository interface {

}

type userrepositories struct {
	db *sqlx.db

}

func Userrepositories(db *sqlx.db) IusersRepository {
	return &userrepositories{
			db : db,
	}
}

