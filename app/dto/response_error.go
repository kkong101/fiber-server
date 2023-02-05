package dto

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprint(e.Message)
}

func Resp(c *fiber.Ctx, resp Error) error {
	return c.JSON(resp)
}

var ErrorHandler = func(c *fiber.Ctx, err error) error {
	resp := Error{
		Code: "9999",
	}

	c.Format()

	if e, ok := err.(validator.ValidationErrors); ok {
		resp.Code = strconv.Itoa(fiber.StatusBadRequest)
		resp.Message = e.Error() + "1111"
	} else if e, ok := err.(*fiber.Error); ok {

		resp.Code = strconv.Itoa(e.Code)
		resp.Message = e.Message + "2222"
	} else if e, ok := err.(*Error); ok {
		resp.Code = strconv.Itoa(fiber.StatusBadRequest)
		resp.Message = e.Message + "3333"

	} else {
		resp.Message = err.Error()
	}
	return Resp(c, resp)
}
