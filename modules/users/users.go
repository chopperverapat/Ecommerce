package users

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Username string `db:"username" json:"username"`
	RoleId   int    `db:"role_id" json:"role_id"`
}

// body request
type UserRegisterReq struct {
	Email    string `db:"email" json:"email" form:"email"`
	Username string `db:"username" json:"username" form:"username"`
	Password string `db:"password" json:"password" form:"password"`
}

func (u *UserRegisterReq) CheckEmail() bool {
	maths, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, u.Email)
	if err != nil {
		fmt.Printf("email is invalid: %v\n", err)
		return false
	}
	return maths
}

func (u *UserRegisterReq) BcrypPass() error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return fmt.Errorf("hashed password failed: %v", err)
	}
	u.Password = string(hashPass)
	return nil
}

type UserPasssport struct {
	user      *User
	usertoken *UserToken
}

type UserToken struct {
	Id           string `json:"id" db:"db"`
	Accesstoken  string `json:"access_token" db:"access_token"`
	Refreshtoken string `json:"refresh_token" db:"refresh_token"`
}
