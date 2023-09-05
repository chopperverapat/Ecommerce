package entities

import (
	"github.com/gofiber/fiber/v2"

	"cpshop/pkg/logger"
)

type IResponse interface {
	Success(code int, data any) IResponse
	Error(code int, traceID, msg string) IResponse
	Res() error
}

type Response struct {
	StatusCode int
	Data       any
	ErrorRes   *ErrorResponse
	Context    *fiber.Ctx
	IsError    bool
}

type ErrorResponse struct {
	TraceId string `json:"trace_id"`
	Msg     string `json:"message"`
}

func NewResponse(c *fiber.Ctx) IResponse {
	return &Response{
		Context: c,
	}

}

func (r *Response) Success(code int, data any) IResponse {
	r.StatusCode = code
	r.Data = data
	// func InitLog รับ res any เลยส่ง &r.Data จะชัวกว่า และ เร็วกว่าส่ง r.Data
	logger.InitLog(r.Context, &r.Data).PrintLog().SaveLog()
	return r
}
func (r *Response) Error(code int, traceID, msg string) IResponse {
	r.StatusCode = code
	r.ErrorRes = &ErrorResponse{
		TraceId: traceID,
		Msg:     msg,
	}
	r.IsError = true
	logger.InitLog(r.Context, &r.ErrorRes).PrintLog().SaveLog()
	return r
}

func (r *Response) Res() error {
	// use any เพราะ type interface เป็นอะไรก็ได้
	return r.Context.Status(r.StatusCode).JSON(func() any {
		if r.IsError {
			// return & เพราะจะได้อ้างอิง address นอก packages
			return &r.ErrorRes
		}
		return &r.Data
	}())
}
