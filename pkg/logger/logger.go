package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"cpshop/pkg/utils"
)

type ILog interface {
	SetQuery(c *fiber.Ctx)
	SetBody(c *fiber.Ctx)
	SetResponse(res any)
	PrintLog() ILog
	SaveLog()
}

type logall struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

// any จะได้ flexible ในการใช้

func InitLog(c *fiber.Ctx, res any, code int) ILog {
	log := &logall{
		Time:       time.Now().Local().Format("2006-01-02 15:04:05"),
		Ip:         c.IP(),
		Method:     c.Method(),
		Path:       c.Path(),
		StatusCode: code,
	}
	log.SetQuery(c)
	log.SetBody(c)
	log.SetResponse(res)
	return log
}

// จริงไม่ต้อง return ตั้งแต่กำหนด struct แค่ว่าใส่มาก่อนละกันเผืิ่อใช้ ค่อยเอาออก
func (l *logall) PrintLog() ILog {
	utils.Debug(l)
	return l
}

func (l *logall) SaveLog() {
	// bytes must to string before write in file log
	logtofile := utils.Outputtofile(l)
	filename := fmt.Sprintf("./assets/logs/log-%v.log", strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""))
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("error open file: %v", err)
	}
	defer file.Close()
	file.WriteString(string(logtofile) + "\n")
}

func (l *logall) SetQuery(c *fiber.Ctx) {
	var query any
	if err := c.QueryParser(&query); err != nil {
		log.Printf("error query parser: %v", err)
	}
	l.Query = query
}
func (l *logall) SetBody(c *fiber.Ctx) {
	var body any
	if err := c.BodyParser(&body); err != nil {
		log.Printf("error body parser: %v", err)
	}
	l.Body = body
	switch l.Path {
	case "/api/v2/signup":
		l.Body = ""
	default:
		l.Body = body
	}
}

func (l *logall) SetResponse(res any) {
	l.Response = res
}

// func (l *logall) SetQuery(c *fiber.Ctx) {
// 	var query any
// 	// var query interface{}
// 	// จริงๆใช้ query ได้ แต่เป็น any ใช้ & ดีกว่า และได้ความเร็วด้วย
// 	if err := c.QueryParser(&query); err != nil {
// 		log.Printf("error query parser: %v", err)
// 	}
// 	l.Query = query
// }

// body เช่น ส่ง user, pass มา
// func (l *logall) SetBody(c *fiber.Ctx) {
// 	var body any
// 	// var body interface{}
// 	if err := c.BodyParser(&body); err != nil {
// 		log.Printf("error body parser: %v", err)
// 	}
// 	l.Body = body
// 	// log จะไม่เก็บ password
// 	switch l.Path {
// 	case "/api/v2/signup":
// 		l.Body = ""
// 	default:
// 		l.Body = body
// 	}
// }

// responese คือสิ่งที่เกิดจาก ฝั่งเราส่งกลับไป ไม่ใช่มาจาก request จีงไม่ต้องใช้ c *fiber.Ctx
// ใช้ res จากที่เราจะส่งกลับไปให้ clients แทน
