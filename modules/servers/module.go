package servers

import (
	"github.com/gofiber/fiber/v2"

	"cpshop/modules/middleware/middlewareHandler"
	"cpshop/modules/middleware/middlewareRepositories"
	"cpshop/modules/middleware/middlewareUsecases"

	"cpshop/modules/monitor/monitorHandler"

	"cpshop/modules/users/userHandler"
	"cpshop/modules/users/userRepositories"
	"cpshop/modules/users/userUsecase"
)

type IModuleFactory interface {
	MonitorModule()
	UsersModules()
}

type moduleFactory struct {
	// route fibe
	r fiber.Router
	// server same package canuser server lowercase
	s   *server
	mid middlewareHandler.IMiddlewareHandler
}

func InitModuleFactory(r fiber.Router, s *server, mid middlewareHandler.IMiddlewareHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewareHandler.IMiddlewareHandler {
	repository := middlewareRepositories.NewMiddlewareRepositories(s.db)
	usecase := middlewareUsecases.NewMiddlewareUsecase(repository)
	return middlewareHandler.NewMiddlewareHandler(s.cfg, usecase)
}

// create router : /api/v1 etc.
func (m *moduleFactory) MonitorModule() {
	handler := monitorHandler.MonitorHandler(m.s.cfg)
	//function call back fiber ไม่ต้องส่ง *fiber,Ctx ไป มัน past มากับ *fiber.Router ละ
	m.r.Get("/", handler.HealthCheck)
}

func (m *moduleFactory) UsersModules() {
	repository := userRepositories.Userrepositories(m.s.db)
	usecase := userUsecase.UserUsecase(m.s.cfg, repository)
	handler := userHandler.Userhandler(m.s.cfg, usecase)

	// group route
	router := m.r.Group("/users")

	// pass auto fiber.Ctx because using handler dont need to use SignupCustome(*fiber.Ctx)
	router.Post("/signup", handler.SignupCustomer) // แก้ไขนี้
}
