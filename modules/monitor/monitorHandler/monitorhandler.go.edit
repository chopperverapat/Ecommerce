package monitorHandler

import (
	"cpshop/config"
	"cpshop/modules/entities"
	"cpshop/modules/monitor"

	"github.com/gofiber/fiber/v2"
)

type ImonitorHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type monitorHanlder struct {
	cfg config.Iconfig
}

func MonitorHandler(cfg config.Iconfig) ImonitorHandler {
	return &monitorHanlder{
		cfg: cfg,
	}
}

func (m *monitorHanlder) HealthCheck(c *fiber.Ctx) error {
	response := &monitor.Monitor{
		Name:    m.cfg.App().Name(),
		Version: m.cfg.App().Version(),
	}
	// return c.Status(fiber.StatusOK).JSON(response)
	return entities.NewResponse(c).Success(fiber.StatusOK, response).Res()
}
