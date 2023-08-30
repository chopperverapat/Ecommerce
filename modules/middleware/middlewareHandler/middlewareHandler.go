package middlewareHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"cpshop/config"
	"cpshop/modules/entities"
	"cpshop/modules/middleware/middlewareUsecases"
)

type middlewareHandlersErrCode string

const (
	routerCheckErr middlewareHandlersErrCode = "middlware-001"
)

type IMiddlewareHandler interface {
	// midleware cors เพื่อให้ทำงานกับแอพอื่นที่ domain,port ต่างกันได้
	Cors() fiber.Handler
	RouterCheck() fiber.Handler
	Logger() fiber.Handler
}

type middlewaresHandler struct {
	cfg               config.Iconfig
	middlewareUsecase middlewareUsecases.IMiddlewareUsecase
}

func NewMiddlewareHandler(cfg config.Iconfig, middlewareUsecase middlewareUsecases.IMiddlewareUsecase) IMiddlewareHandler {
	return &middlewaresHandler{
		cfg:               cfg,
		middlewareUsecase: middlewareUsecase,
	}
}

func (m *middlewaresHandler) Cors() fiber.Handler {
	// build in cors middleware เลยใช้ cors ได้
	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}

func (h *middlewaresHandler) RouterCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		response := entities.NewResponse(c)
		return response.Error(
			fiber.ErrNotFound.Code,
			string(routerCheckErr),
			"router not found",
		).Res()
	}
}

func (h *middlewaresHandler) Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${time} [${ip}] ${status} - ${method} ${path}\n",
		TimeFormat: "02/01/2006",
		TimeZone:   "Bangkok/Asia",
	})
}
