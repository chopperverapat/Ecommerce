package patterns

import (
	"context"
	"encoding/json"
	"fmt"
	"server/modules/users"
	"time"
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
	db      *sqlx.db
}

type admin struct {
	*userrequest
}

type customer struct {
	*userrequest
}

func InsertUsers(db *sqlx.db, body *users.UserRegisterReq, isAdmin bool) Iinsertuser {
	if isAdmin {
		return newadmin(db, body)
	}
	return newCustomer(db, body)

}

func newCustomer(db *sqlx.db, body *users.UserRegisterRe) Iinsertuser {
	return &customer{
		userrequest: &userrequest{
			bodyreq: body,
			db:      db,
		},
	}
}

func newadmin(db *sqlx.db, body *users.UserRegisterRe) Iinsertuser {
	return &admin{
		userrequest: &userrequest{
			bodyreq: body,
			db:      db,
		},
	}
}
func (uq *userrequest) Customer() (Iinsertuser, error) {
	// wait time
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	// if failed retrun cancle and give resresource to server
	defer cancle()
	// assume role_id 1
	query := `
	INSERT NTO "users"(
		"username"
		"email"
		"password"
		"role_id"
		)
	VALUES
		($1, $2, $3, 1)
	RETURNING "id";
	`
	if err := uq.db.QueryRowContext(
		ctx,
		query,
		uq.bodyreq.Username,
		uq.bodyreq.Email,
		uq.bodyreq.Password,
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
	return nil, nil
}

func (uq *userrequest) Admin() (Iinsertuser, error) {
	return nil, nil
}
func (uq *userrequest) Result() (*users.UserPasssport, error) {
	// query db to json after pass to struct
	query := `
	SELECT 
			json_build_boj(
				'user': "t",
				'token': NULL
			)
	FROM (
		SELECT 
				"u"."id"
				"u"."username"
				"u"."password"
				"u"."role_id"
		FROM users "u"
		HERE "u"."id" = $1
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
	return user, nil
}
