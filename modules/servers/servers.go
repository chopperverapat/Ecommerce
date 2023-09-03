package servers

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"cpshop/config"
)

type server struct {
	app *fiber.App
	cfg config.Iconfig
	db  *sqlx.DB
}

func Newservers(cfg config.Iconfig, db *sqlx.DB) Iserver {
	return &server{
		app: fiber.New(fiber.Config{
			AppName:      cfg.App().Name(),
			BodyLimit:    cfg.App().BodyLimit(),
			ReadTimeout:  cfg.App().ReadTimeout(),
			WriteTimeout: cfg.App().WriteTimeout(),
			JSONEncoder:  json.Marshal,   // make fiber fast
			JSONDecoder:  json.Unmarshal, // make fiber fast ref doc fiber
		}),
		cfg: cfg,
		db:  db,
	}
}

type Iserver interface {
	// method return
	// no return
	Start()
}

func (s *server) Start() {
	// init middleware
	mid := InitMiddlewares(s)
	s.app.Use(mid.Logger())
	s.app.Use(mid.Cors())

	v1 := s.app.Group("/api/v2")
	// s ต้องส่ง * ไปให้ InitModule แต่เราประกาศ s *server ไปแล้ว จึงใช้ s ได้เลย
	modules := InitModuleFactory(v1, s, mid)
	modules.MonitorModule()
	modules.UsersModules()

	s.app.Use(mid.RouterCheck())

	// graceful shutdown ใช้ในการ เคลียร์ระบบ ก่อนจะปิดแบบคลีนๆ ในกรณี ล่ม,สัญญาณมารบกวน,บลาๆๆๆ
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// go concurrency
	go func() {
		_ = <-c
		log.Println("Server is shutting down...")
		// method fiber : turn off app
		_ = s.app.Shutdown()

	}()

	log.Printf("server is starting on %v", s.cfg.App().Url())
	s.app.Listen(s.cfg.App().Url())

}
