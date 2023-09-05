package patterns

import (
	"context"
	"cpshop/modules/users"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

//factory to seperate admin, customer

type Iinsertuser interface {
	// return for use method result
	Customer() (Iinsertuser, error)
	Admin() (Iinsertuser, error)
	Result() (*users.UserPasssport, error)
}

type userrequest struct {
	id      string
	bodyreq *users.UserRegisterReq
	db      *sqlx.DB
}

type admin struct {
	*userrequest
}

type customer struct {
	*userrequest
}

func InsertUsers(db *sqlx.DB, body *users.UserRegisterReq, isAdmin bool) Iinsertuser {
	if isAdmin {
		return newadmin(db, body)
	}
	return newCustomer(db, body)

}

func newCustomer(db *sqlx.DB, body *users.UserRegisterReq) Iinsertuser {
	return &customer{
		userrequest: &userrequest{
			bodyreq: body,
			db:      db,
		},
	}
}

func newadmin(db *sqlx.DB, body *users.UserRegisterReq) Iinsertuser {
	return &admin{
		userrequest: &userrequest{
			bodyreq: body,
			db:      db,
		},
	}
}
func (uq *userrequest) Customer() (Iinsertuser, error) {
	// wait time
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// if failed retrun cancle and give resresource to server
	defer cancel()

	// assume role_id 1
	query := `
	INSERT INTO "users" (
		"email",
		"password",
		"username",
		"role_id"
	)
	VALUES
		($1, $2, $3, 2)
	RETURNING "id";`
	if err := uq.db.QueryRowContext(
		ctx,
		query,
		uq.bodyreq.Email,
		uq.bodyreq.Password,
		uq.bodyreq.Username,
	// scan use with sql RETURNING
	// return id db to userrequest struct
	).Scan(&uq.id); err != nil {
		switch err.Error() {
		case "ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)":
			return nil, fmt.Errorf("username has been used")
		case "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)":
			return nil, fmt.Errorf("email has been used")
		default:
			return nil, fmt.Errorf("insert user failed: %v", err)
		}
	}
	// fmt.Printf("user_insert: %v,\n", uq)
	return uq, nil
}

func (uq *userrequest) Admin() (Iinsertuser, error) {
	return nil, nil
}
func (uq *userrequest) Result() (*users.UserPasssport, error) {
	// query db to json after pass to struct
	if uq != nil && uq.bodyreq != nil {
		query := `
		SELECT
			json_build_object(
				'user', "t",
				'token', NULL
			)
		FROM (
			SELECT
				"u"."id",
				"u"."email",
				"u"."username",
				"u"."role_id"
			FROM "users" "u"
			WHERE "u"."id" = $1
		) AS "t"`

		// query json bytes
		data := make([]byte, 0)
		//Get = query 1 row
		//Select = query multi rows

		//data to recieve query
		if err := uq.db.Get(&data, query, uq.id); err != nil {
			return nil, fmt.Errorf("get user failed: %v", err)
		}
		// create new struct to recieve byte=> json to struct
		user := new(users.UserPasssport)
		// data to struct user by unmarshal
		if err := json.Unmarshal(data, &user); err != nil {
			return nil, fmt.Errorf("unmarshal user failed: %v", err)
		}
		// fmt.Printf("user_select: %v,\n", user)
		return user, nil
	}
	return nil, fmt.Errorf("UserRequest is nil or isAdmin is true")
}
