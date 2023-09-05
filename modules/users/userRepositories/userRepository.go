package userRepositories

import (
	"cpshop/modules/users"
	"errors"

	"cpshop/modules/users/patterns"

	"github.com/jmoiron/sqlx"
)

type IusersRepository interface {
	InsertUsers(uq *users.UserRegisterReq, isAdmin bool) (*users.UserPasssport, error)
}

type userrepositories struct {
	db *sqlx.DB
}

func Userrepositories(db *sqlx.DB) IusersRepository {
	return &userrepositories{
		db: db,
	}
}

func (ur *userrepositories) InsertUsers(uq *users.UserRegisterReq, isAdmin bool) (*users.UserPasssport, error) {
	result := patterns.InsertUsers(ur.db, uq, isAdmin)
	// fmt.Printf("RESULT: %v", result)
	// Check if result is nil
	if result == nil {
		return nil, errors.New("patterns.InsertUsers returned nil result")
	}
	// fmt.Printf("Result from patterns.InsertUsers: %#v\n", result)
	var err error
	if isAdmin {
		result, err = result.Admin()
		if err != nil {
			return nil, err
		}
	} else {
		result, err = result.Customer()
		if err != nil {
			return nil, err
		}
	}

	// get result from insert
	userresult, err := result.Result()
	if err != nil {
		return nil, err
	}
	// fmt.Printf("userresult_repo: %v,\n", userresult)
	return userresult, nil
}
