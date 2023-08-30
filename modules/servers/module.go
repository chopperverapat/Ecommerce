package servers

import (
	"github.com/gofiber/fiber/v2"

	"cpshop/modules/middleware/middlewareHandler"
	"cpshop/modules/middleware/middlewareRepositories"
	"cpshop/modules/middleware/middlewareUsecases"
	"cpshop/modules/monitor/monitorHandler"
)

type IModuleFactory interface {
	MonitorModule()
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
