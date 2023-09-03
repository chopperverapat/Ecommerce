package userHandler

import (
	"cpshop/config"
	"cpshop/modules/entities"
	"cpshop/modules/users"
	"cpshop/modules/users/userUsecase"

	"github.com/gofiber/fiber/v2"
)

//using constant Error

type ErruserHandler string

const (
	signupCustomer ErruserHandler = "users-001"
)

type IuserHandler interface {
	SignupCustomer(c *fiber.Ctx) error
}

type userhandler struct {
	cfg         config.Iconfig
	userUsecase userUsecase.IuserUsercase
}

func Userhandler(cfg config.Iconfig, userUsecase userUsecase.IuserUsercase) IuserHandler {
	return &userhandler{
		cfg:         cfg,
		userUsecase: userUsecase,
	}
}

func (uh *userhandler) SignupCustomer(c *fiber.Ctx) error {
	// new create pointer
	reqBody := new(users.UserRegisterReq)
	// reqBody pointer can pass , dont need user &
	if err := c.BodyParser(reqBody); err != nil {
		// if err := c.BodyParser(reqBody); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signupCustomer),
			err.Error(),
		).Res()
	}
	// Email check
	if !reqBody.CheckEmail() {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signupCustomer),
			"Email is invalid",
		).Res()
	}

	// insert call suecase
	insertuser, err := uh.userUsecase.InsertCustomer(reqBody)
	if err != nil {
		switch err.Error() {
		case "username has been used":
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(signupCustomer),
				err.Error(),
			).Res()
		case "email has been used":
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(signupCustomer),
				err.Error(),
			).Res()
		default:
			return entities.NewResponse(c).Error(
				fiber.ErrInternalServerError.Code,
				string(signupCustomer),
				err.Error(),
			).Res()
		}
	}
	return entities.NewResponse(c).Success(fiber.StatusOK, insertuser).Res()
}
